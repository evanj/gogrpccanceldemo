# Go gRPC Deadlines and Cancelation Demo

An example of deadlines and cancellation with a Go gRPC server and client. It sends the following requests:

* Plain request that returns normally. gRPC `OK`.
* Client times out. gRPC `DEADLINE_EXCEEDED`.
* Server returns a Go `context.DeadlineExceeded` error. This simulates what happens when the server hits a timeout while handling the request and returns it. gRPC `DEADLINE_EXCEEDED`.
* Server does parallel work, one branch of which times out, causing the other to be canceled. This is a "bug" in server code, but does happen. gRPC `CANCELED`.
* Client cancels the request. Client-side gRPC `CANCELED`. Server can check ctx.Err() to detect.
* Client makes a gRPC call with an already cancelled context (NO SERVER REQUEST). Client-side gRPC `CANCELED`.
* Client makes a gRPC call with an already timeed out context (NO SERVER REQUEST). Client-side gRPC `DEADLINE_EXCEEDED`.

This helped me understand how the client and server handle these scenarios.


## Observations

* The Go gRPC server converts `context.DeadlineExceeded` and `context.Canceled` return values into a gRPC error status `DEADLINE_EXCEEDED` and `CANCELED` (respectively).
* Go gRPC client always returns a gRPC status error, even if the client-side context is canceled or timed out. It does not make any requests in these cases.


## Detecting a canceled request on the server

The incoming `ctx.Err() == context.Canceled`. If this is a top level handler without any interceptors that have modified the context, this should be the only way for this to occur. An interceptor could in theory cancel the context itself, but that would be unusual.


## Distinguishing between client-side and server-side timeout/cancels on the client

Since the return value from a gRPC client call is always a gRPC status object, you need to examine the `ctx.Err()`.

* `ctx.Err() == nil`: Must be a server-side result.
* `ctx.Err() == context.DeadlineExceeded && status.Convert(err).Code() == codes.DeadlineExceeded`: Most likely a client-side timeout.
* `ctx.Err() == context.Canceled && status.Convert(err).Code() == codes.Canceled`: Most likely a client-side cancel.

Note that this is racy, so it is not guaranteed it is a client-side error. It is possible for the server to return `DeadlineExceeded` error, the client to receive and return that, *then* for the Go context to time out. In this situation, it shouldn't matter: both happened at close enough to the same time.


## Usage

1. Start the server: `go run ./echoserver`
2. Run the client: `go run ./echoclient`

The client has some flags to control some features which can be shown with `--help`.

To regenerate the proto files, run `make`.


## Example server output

* reqID=2 is the timed out request and has `ctx.Err() == context.DeadlineExceeded`
* reqID=5 is the canceled request and has `ctx.Err() == context.Canceled`

```
2024/02/05 17:10:16 serving on localhost:8001 ...
2024/02/05 17:10:20 reqID=1 starting with deadline? false 0001-01-01T00:00:00Z duration = 0s
2024/02/05 17:10:20 reqID=1 echoing 26 bytes after sleeping 0s; ctx.Err()=<nil>
2024/02/05 17:10:20 reqID=2 starting with deadline? true 2024-02-05T17:10:20.53399225-05:00 duration = 99.910667ms
2024/02/05 17:10:20 reqID=3 starting with deadline? false 0001-01-01T00:00:00Z duration = 0s
2024/02/05 17:10:20 reqID=3 Action=RETURN_CONTEXT_DEADLINE_EXCEEDED; returning err=context.deadlineExceededError context deadline exceeded; == context.DeadlineExceeded? true
2024/02/05 17:10:20 reqID=4 starting with deadline? false 0001-01-01T00:00:00Z duration = 0s
2024/02/05 17:10:20 reqID=4 Action=RETURN_CONTEXT_CANCELED; spawning two sub-tasks ...
2024/02/05 17:10:20   sub task one returning err=context.deadlineExceededError context deadline exceeded
2024/02/05 17:10:20   parent task collected err=context.deadlineExceededError context deadline exceeded; canceling sub-tasks
2024/02/05 17:10:20   sub task two returning err=*errors.errorString context canceled
2024/02/05 17:10:20   parent task collected err=*errors.errorString context canceled; canceling sub-tasks
2024/02/05 17:10:20   parent task returning err=*errors.errorString context canceled
2024/02/05 17:10:20 reqID=4 Action=RETURN_CONTEXT_CANCELED; returning err=*errors.errorString context canceled; == context.Canceled? true
2024/02/05 17:10:20 reqID=5 starting with deadline? false 0001-01-01T00:00:00Z duration = 0s
2024/02/05 17:10:21 reqID=2 echoing 26 bytes after sleeping 1s; ctx.Err()=context.deadlineExceededError{}
2024/02/05 17:10:21 reqID=5 echoing 26 bytes after sleeping 1s; ctx.Err()=&errors.errorString{s:"context canceled"}
```

## Example client output

```
2024/02/05 17:10:20 setting Dial option WithBlock
2024/02/05 17:10:20 setting Dial timeout=1m0s
2024/02/05 17:10:20 connecting to addr=localhost:8001 ...
2024/02/05 17:10:20 plain request ...
2024/02/05 17:10:20   resp=output:"echoed: example echo request input"
2024/02/05 17:10:20 client timeout shorter than the server-side sleep ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=DeadlineExceeded Message="context deadline exceeded" Details=[]interface {}{}
2024/02/05 17:10:20 server returns Go context deadline exceeded ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=DeadlineExceeded Message="context deadline exceeded" Details=[]interface {}{}
2024/02/05 17:10:20 server returns Go context canceled ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=Canceled Message="context canceled" Details=[]interface {}{}
2024/02/05 17:10:20 client cancels request ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=Canceled Message="context canceled" Details=[]interface {}{}
2024/02/05 17:10:20 client request with already cancelled context (NO SERVER REQUEST) ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=Canceled Message="context canceled" Details=[]interface {}{}
2024/02/05 17:10:20 client request with already timed out context (NO SERVER REQUEST) ...
2024/02/05 17:10:20   resp=(*echopb.EchoResponse)(nil) grpcErr.Code=DeadlineExceeded Message="context deadline exceeded" Details=[]interface {}{}
```
