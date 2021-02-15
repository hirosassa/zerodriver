package zerodriver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

// HTTPPayload is the struct consists of http request related components.
// Details are in following link.
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
type HTTPPayload struct {
	RequestMethod                  string  `json:"requestMethod"`
	RequestURL                     string  `json:"requestUrl"`
	RequestSize                    string  `json:"requestSize"`
	Status                         int     `json:"status"`
	ResponseSize                   string  `json:"responseSize"`
	UserAgent                      string  `json:"userAgent"`
	RemoteIP                       string  `json:"remoteIp"`
	ServerIP                       string  `json:"serverIp"`
	Referer                        string  `json:"referer"`
	Latency                        Latency `json:"latency"`
	CacheLookup                    bool    `json:"cacheLookup"`
	CacheHit                       bool    `json:"cacheHit"`
	CacheValidatedWithOriginServer bool    `json:"cacheValidatedWithOriginServer"`
	CacheFillBytes                 string  `json:"cacheFillBytes"`
	Protocol                       string  `json:"protocol"`
}

// The request processing latency on the server, from the time the request was
// received until the response was sent.
type Latency struct {
	Nanos   int32 `json:"nanos"`
	Seconds int64 `json:"seconds"`
}

func (e *Event) HTTP(req *HTTPPayload) *zerolog.Event {
	return e.Event.Interface("httpRequest", req)
}

// NewHTTP returns a HTTPPayload struct.
func NewHTTP(req *http.Request, res *http.Response) *HTTPPayload {
	if req == nil {
		req = &http.Request{}
	}

	if res == nil {
		res = &http.Response{}
	}

	payload := &HTTPPayload{
		RequestMethod: req.Method,
		RequestSize:   strconv.FormatInt(req.ContentLength, 10),
		Status:        res.StatusCode,
		ResponseSize:  strconv.FormatInt(res.ContentLength, 10),
		UserAgent:     req.UserAgent(),
		RemoteIP:      req.RemoteAddr,
		Referer:       req.Referer(),
		Protocol:      req.Proto,
	}

	if req.URL != nil {
		payload.RequestURL = req.URL.String()
	}

	return payload
}

// MakeLatency returns Latency struct based on passed time.Duration object.
func MakeLatency(d time.Duration) Latency {
	nanos := d.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return Latency{
		Nanos:   int32(nanos),
		Seconds: secs,
	}
}
