# zerodriver

[![Actions Status: test](https://github.com/hirosassa/zerodriver/workflows/test/badge.svg)](https://github.com/hirosassa/zerodriver/actions?query=workflow%3A"test")
[![Actions Status: golangci-lint](https://github.com/hirosassa/zerodriver/workflows/golangci-lint/badge.svg)](https://github.com/hirosassa/zerodriver/actions?query=workflow%3A"golangci-lint")
[![Apache-2.0](https://img.shields.io/github/license/hirosassa/zerodriver)](LICENSE)


[Zerolog](https://github.com/rs/zerolog) based [Cloud Logging](https://cloud.google.com/logging) (formerly Stackdriver Logging). This package is inspired by [Zapdriver](https://github.com/blendle/zapdriver).

## What is this package?

This package provides simple structured logger optimized for [Cloud Logging](https://cloud.google.com/logging) based on [zerolog](https://github.com/rs/zerolog).

Key features of tbls are:

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
    logger := zerodriver.New(false)
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
func (e *Event) HTTP(req *zapdriver.HTTPPayload) *zerolog.Event
```

This function feature is borrowed from zapdriver. You can generate `zapdriver.HTTPPayload` from `http.Request` and `http.Response` using `NewHTTP` function.
Detail description of usage of these function is available [here](https://github.com/blendle/zapdriver#http).
Using these feature, you can log HTTP related information as follows,

```go
p := NewHTTP(req, res)
logger.Info().HTTP(p).Msg("request received")
```
