package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/evanj/gogrpccanceldemo/echopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	addr := flag.String("addr", "localhost:8001", "server address")
	withBlock := flag.Bool("withBlock", true, "if we should use the WithBlock dial option")
	dialTimeout := flag.Duration("dialTimeout", time.Minute, "timeout to use for DialContext")
	input := flag.String("input", "example echo request input", "message to echo")
	flag.Parse()

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if *withBlock {
		log.Printf("setting Dial option WithBlock")
		dialOptions = append(dialOptions, grpc.WithBlock())
	}

	dialCtx := context.Background()
	cancel := func() {}
	if *dialTimeout > 0 {
		dialCtx, cancel = context.WithTimeout(dialCtx, *dialTimeout)
		log.Printf("setting Dial timeout=%s", dialTimeout.String())
	}
	log.Printf("connecting to addr=%s ...", *addr)
	conn, err := grpc.DialContext(dialCtx, *addr, dialOptions...)
	cancel()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := echopb.NewEchoClient(conn)

	// normal request
	log.Printf("plain request ...")
	ctx := context.Background()
	resp, err := client.Echo(ctx, &echopb.EchoRequest{Input: *input})
	if err != nil {
		panic(err)
	}
	log.Printf("  resp=%s", resp.String())

	// client timeout
	log.Printf("client timeout shorter than the server-side sleep ...")
	const clientTimeout = 100 * time.Millisecond
	ctxTimeout, cancel := context.WithTimeout(ctx, clientTimeout)
	req := &echopb.EchoRequest{
		Input:       *input,
		ServerSleep: durationpb.New(10 * clientTimeout),
	}
	resp, err = client.Echo(ctxTimeout, req)
	cancel()
	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.DeadlineExceeded {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected deadline exceeded error; resp=%#v err=%s", resp, err))
	}

	// server timeout
	log.Printf("server returns Go context deadline exceeded ...")
	req = &echopb.EchoRequest{
		Input:  *input,
		Action: echopb.ServerAction_RETURN_CONTEXT_DEADLINE_EXCEEDED,
	}
	resp, err = client.Echo(ctx, req)
	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.DeadlineExceeded {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected deadline exceeded error; resp=%#v err=%s", resp, err))
	}

	// server returns context canceled
	log.Printf("server returns Go context canceled ...")
	req = &echopb.EchoRequest{
		Input:  *input,
		Action: echopb.ServerAction_RETURN_CONTEXT_CANCELED,
	}
	resp, err = client.Echo(ctx, req)
	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.Canceled {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected canceled error; resp=%#v err=%s", resp, err))
	}

	// client cancels the request after a short wait, sending a cancel to the server
	log.Printf("client cancels request ...")
	req = &echopb.EchoRequest{
		Input:       *input,
		ServerSleep: durationpb.New(time.Second),
	}
	ctxWithCancel, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		resp, err = client.Echo(ctxWithCancel, req)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		cancel()
	}()
	wg.Wait()

	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.Canceled {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected canceled error; resp=%#v err=%s", resp, err))
	}

	// context already canceled
	log.Printf("client request with already cancelled context (NO SERVER REQUEST) ...")
	if ctxWithCancel.Err() != context.Canceled {
		panic("expected context to be already canceled")
	}
	req = &echopb.EchoRequest{
		Input: *input,
	}
	resp, err = client.Echo(ctxWithCancel, req)
	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.Canceled {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected canceled error; resp=%#v err=%s", resp, err))
	}

	// context already timed out
	log.Printf("client request with already timed out context (NO SERVER REQUEST) ...")
	if ctxTimeout.Err() != context.DeadlineExceeded {
		panic("expected context to be already timed out")
	}
	req = &echopb.EchoRequest{
		Input: *input,
	}
	resp, err = client.Echo(ctxTimeout, req)
	if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.DeadlineExceeded {
		log.Printf("  resp=%#v grpcErr.Code=%s Message=%#v Details=%#v",
			resp, grpcErr.Code().String(), grpcErr.Message(), grpcErr.Details())
	} else {
		panic(fmt.Sprintf("expected deadline exceeded error; resp=%#v err=%s", resp, err))
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
