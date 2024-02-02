package main

import (
	"context"
	"flag"
	"log"
	"net"
	"sync/atomic"
	"time"

	"github.com/evanj/gogrpccanceldemo/echopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	echopb.UnimplementedEchoServer
	responseSleep time.Duration
	requestID     atomic.Int64
}

func newServer(responseSleep time.Duration) *server {
	return &server{echopb.UnimplementedEchoServer{}, responseSleep, atomic.Int64{}}
}

// simTwoTasks simulates running two tasks in parallel that respect cancelation. One task times out,
// which cases the other to be canceled. The parent task's error handling returns the last error
// instead of the first, so it return context.Canceled instead of context.DeadlineExceeded.
// This is arguably a bug: the code should return the deadline exceeded error since that happened
// first. However, there are many ways to write code that does something like this.
func simTwoTasks(ctx context.Context) error {
	collectErrs := make(chan error)

	// create a context that gets cancelled if an error occurs in either task
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	const numTasks = 2

	go func() {
		// task one: times out
		ctx, cancel := context.WithTimeout(ctxWithCancel, 10*time.Nanosecond)
		<-ctx.Done()
		cancel()
		err := ctx.Err()
		log.Printf("  sub task one returning err=%T %s", err, err)
		collectErrs <- err
	}()

	go func() {
		// task two: simulates a longer task that respects cancelation (e.g. a gRPC call)
		<-ctxWithCancel.Done()
		err := ctxWithCancel.Err()
		log.Printf("  sub task two returning err=%T %s", err, err)
		collectErrs <- err
	}()

	// collect errors from the sub-tasks: this "transforms" the deadline exceeded into canceled
	var lastErr error
	for i := 0; i < numTasks; i++ {
		lastErr = <-collectErrs
		if lastErr != nil {
			log.Printf("  parent task collected err=%T %s; canceling sub-tasks", lastErr, lastErr)
			cancel()
		}
	}
	log.Printf("  parent task returning err=%T %s", lastErr, lastErr)
	return lastErr
}

func (s *server) Echo(ctx context.Context, request *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	requestID := s.requestID.Add(1)

	deadline, ok := ctx.Deadline()
	deadlineDuration := 0 * time.Second
	if ok {
		deadlineDuration = time.Until(deadline)
	}
	log.Printf("reqID=%d starting with deadline? %t %s duration = %s",
		requestID, ok, deadline.Format(time.RFC3339Nano), deadlineDuration.String())

	// handle a special action
	switch request.Action {
	case echopb.ServerAction_UNSPECIFIED:
		// no special action: normal handling below

	case echopb.ServerAction_RETURN_CONTEXT_DEADLINE_EXCEEDED:
		ctx, cancel := context.WithTimeout(ctx, 10*time.Nanosecond)
		<-ctx.Done()
		cancel()
		err := ctx.Err()
		log.Printf("reqID=%d Action=%s; returning err=%T %s; == context.DeadlineExceeded? %t",
			requestID, request.Action.String(), err, err, err == context.DeadlineExceeded)
		return nil, err

	case echopb.ServerAction_RETURN_CONTEXT_CANCELED:
		log.Printf("reqID=%d Action=%s; spawning two sub-tasks ...",
			requestID, request.Action.String())
		err := simTwoTasks(ctx)
		log.Printf("reqID=%d Action=%s; returning err=%T %s; == context.Canceled? %t",
			requestID, request.Action, err, err, err == context.Canceled)
		return nil, err

	default:
		log.Printf("reqID=%d Unknown Action=%s %d; returning gRPC error",
			requestID, request.Action.String(), request.Action.Number())
		return nil, status.Errorf(codes.InvalidArgument, "unknown action value %d", request.Action.Number())
	}

	sleep := s.responseSleep
	requestSleep := request.ServerSleep.AsDuration()
	if sleep < requestSleep {
		sleep = requestSleep
	}
	time.Sleep(sleep)
	log.Printf("reqID=%d echoing %d bytes after sleeping %s; ctx.Err()=%#v",
		requestID, len(request.Input), sleep.String(), ctx.Err())

	resp := &echopb.EchoResponse{
		Output: "echoed: " + request.Input,
	}
	return resp, nil
}

func main() {
	addr := flag.String("addr", "localhost:8001", "listening address")
	responseSleep := flag.Duration("responseSleep", 0, "time to sleep before responding")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	echopb.RegisterEchoServer(s, newServer(*responseSleep))

	log.Printf("serving on %s ...", *addr)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
