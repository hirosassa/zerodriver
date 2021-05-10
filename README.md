# zerodriver

[![Actions Status: test](https://github.com/hirosassa/zerodriver/workflows/test/badge.svg)](https://github.com/hirosassa/zerodriver/actions?query=workflow%3A"test")
[![Actions Status: golangci-lint](https://github.com/hirosassa/zerodriver/workflows/golangci-lint/badge.svg)](https://github.com/hirosassa/zerodriver/actions?query=workflow%3A"golangci-lint")
[![Go Reference](https://pkg.go.dev/badge/github.com/hirosassa/zerodriver.svg)](https://pkg.go.dev/github.com/hirosassa/zerodriver)
[![Go Report Card](https://goreportcard.com/badge/github.com/hirosassa/zerodriver)](https://goreportcard.com/report/github.com/hirosassa/zerodriver)
[![Coverage Status](https://coveralls.io/repos/github/hirosassa/zerodriver/badge.svg?branch=master)](https://coveralls.io/github/hirosassa/zerodriver?branch=master)
[![Apache-2.0](https://img.shields.io/github/license/hirosassa/zerodriver)](LICENSE)

[Zerolog](https://github.com/rs/zerolog) based [Cloud Logging](https://cloud.google.com/logging) (formerly Stackdriver Logging). This package is inspired by [Zapdriver](https://github.com/blendle/zapdriver).

## What is this package?

This package provides simple structured logger optimized for [Cloud Logging](https://cloud.google.com/logging) based on [zerolog](https://github.com/rs/zerolog).

Key features of zerodriver are:

- zerolog based simple method chaining API
- optimized for [Cloud Logging LogEntry](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry) format

## Usage

First of all, initialize a logger.

```go
logger := zerodriver.NewProductionLogger() // production mode (global log level set to `info`)
logger := zerodriver.NewDevelopmentLogger() // development mode (global log level set to `debug`)
```

Then, write logs by using zerolog based fluent API!
```go
logger.Info().Str("key", "value").Msg("Hello World!")
// output: {"severity":"INFO","key":"value","time":"2009-11-10T23:00:00Z","message":"hello world"}
```

Here's complete example:

```go
package main

import (
    "github.com/hirosassa/zerodriver"
)

func main() {
    logger := zerodriver.NewProductionLogger()
    logger.Info().Str("key", "value").Msg("hello world")
}

// output: {"severity":"INFO","key":"value","time":"2009-11-10T23:00:00Z","message":"hello world"}
```

### GCP specific fields

If your log follows [LogEntry](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry) format,
you can query logs or create metrics alert easier and efficiently on GCP Cloud Logging console.

#### HTTP request

To log HTTP related metrics and information, you can use following function

```go
func (e *Event) HTTP(req *HTTPPayload) *zerolog.Event
```

This feature is forked from zapdriver. You can generate `zerodriver.HTTPPayload` from `http.Request` and `http.Response` using `NewHTTP` function.
Same as zapdriver.NewHTTP, following fields needs to be set manually:

- `ServerIP`
- `Latency`
- `CacheLookup`
- `CacheHit`
- `CacheValidatedWithOriginServer`
- `CacheFillBytes`

Using these feature, you can log HTTP related information as follows,

```go
p := NewHTTP(req, res)
p.Latency = time.Since(start) // add some fields manually
logger.Info().HTTP(p).Msg("request received")
```

#### Trace context

To add trace information to your log, you can use `TraceContext`. The signature of the function is as follows:
```go
func (e *Event) TraceContext(trace string, spanId string, sampled bool, projectID string) *zerolog.Event
```

You can use this feature as follows:

```go
import	"go.opencensus.io/trace"

span := trace.FromContext(r.Context()).SpanContext()
logger.Info().TraceContext(span.TraceID.String(), span.SpanID.String(), true, "my-project").Msg("trace contexts")

// {"severity":"INFO","logging.googleapis.com/trace":"projects/my-project/traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000","logging.googleapis.com/trace_sampled":true,"message":"trace contexts"}
```

#### Labels

You can add any "labels" to your log by following:

```go
logger.Info().Labels(zerodriver.Label("foo", "var")).Msg("labeled log")

// {"severity":"INFO","logging.googleapis.com/labels":{"foo":"var"},"message":"labeled log"}
```

#### Operations

You can add additional information about a potentially long-running operation with which a log entry is associated by following function:

```go
func (e *Event) Operation(id, producer string, first, last bool) *zerolog.Event
```
Log entries with the same `id` are assumed to be part of the same operation.
The producer is an arbitrary identifier that should be globally unique amongst all the logs of all your applications (meaning it should probably be the unique name of the current application).
You should set `first` to true for the first log in the operation, and `last` to true for the final log of the operation.

Also see, https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogEntryOperation

For readable implementation of `operation` log, you can use following functions:

```go
func (e *Event) OperationStart(id, producer string) *zerolog.Event
func (e *Event) OperationContinue(id, producer string) *zerolog.Event
func (e *Event) OperationEnd(id, producer string) *zerolog.Event
```

A concrete example of operation log is as follows:

```go
logger.Info().OperationStart("foo", "bar").Msg("started")
logger.Debug().OperationContinue("foo", "bar").Msg("processing")
logger.Info().OperationEnd("foo", "bar").Msg("done")
```
