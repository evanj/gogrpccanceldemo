# Go gRPC Deadlines and Cancelation Demo

An example of deadlines and cancellation with a Go gRPC server and client. It sends the following requests:

* Plain request that returns normally. gRPC `OK`.
* Client times out. gRPC `DEADLINE_EXCEEDED`.
* Server returns a Go native context deadline exceeded error. This simulates what happens when the server hits a timeout internally while handling the request and returns it. gRPC `DEADLINE_EXCEEDED`.
* Server does parallel work, one branch of which times out, causing the other to be canceled. gRPC `CANCELED`.
* Client cancels the request on the server. Client-side gRPC `CANCELED`; Server can detect this.
* Client makes a gRPC call with an already cancelled context (NO SERVER REQUEST). Client-side gRPC `CANCELED`.
* Client makes a gRPC call with an already timeed out context (NO SERVER REQUEST). Client-side gRPC `DEADLINE_EXCEEDED`.

This is useful for understanding how the client and server handle these different scenarios.


## Observations

* The Go gRPC server converts `context.DeadlineExceeded` and `context.Canceled` return values into a gRPC error status `DEADLINE_EXCEEDED` and `CANCELED` (respectively).
* Go gRPC client always returns a gRPC status error, even if the client-side context is canceled or timed out.


## Detecting a canceled request on the server

The incoming `ctx.Err() == context.Canceled`. If this is a top level handler without any interceptors that have modified the context, this should be the only way for this to occur.


## Distinguishing between client-side and server-side time out/cancels

Since the return value from a gRPC client call is always a gRPC status object, you need to examine the `ctx.Err()`.

* `ctx.Err() == nil`: Must be a server-side result.
* `ctx.Err() == context.DeadlineExceeded && status.Convert(err).Code() == codes.DeadlineExceeded`: Most likely a client-side timeout.
* `ctx.Err() == context.Canceled && status.Convert(err).Code() == codes.Canceled`: Most likely a client-side cancel.

Note that this is racy, so it is not guaranteed it is a client-side error only. It is possible but unlikely for the server to return a DeadlineExceeded error, the client to receive and return that, *then* for the Go context to time out. In this situation, it shouldn't matter: both happened at close enough to the same time.


## Usage

1. Start the server: `go run ./echoserver`
2. Run the client: `go run ./echoclient`

The client has some flags to control some features which can be shown with `--help`.

To regenerate the proto files, run `make`.
