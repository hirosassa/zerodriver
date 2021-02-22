package zerodriver

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

// HTTPPayloadGKE is the struct consists of http request related components.
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
	CacheFillBytes                 string  `json:"cacheFillByte"`
	Protocol                       string  `json:"protocol"`
}

// Latency is the interface of the request processing latency on the server.
// The format of the Latency should differ for GKE and for GAE, Cloud Run.
type Latency interface{}

// GAELatency is the Latency for GAE and Cloud Run.
type GAELatency struct {
	Seconds int64
	Nanos   int32
}

// HTTP adds thehttpRequest field to the *zerolog.Event context
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
		Status:        res.StatusCode,
		UserAgent:     req.UserAgent(),
		RemoteIP:      remoteIP(req),
		Referer:       req.Referer(),
		Protocol:      req.Proto,
	}

	if req.URL != nil {
		payload.RequestURL = req.URL.String()
	}

	if req.Body != nil {
		payload.RequestSize = strconv.FormatInt(req.ContentLength, 10)
	}

	if res.Body != nil {
		payload.ResponseSize = strconv.FormatInt(res.ContentLength, 10)
	}

	return payload
}

// MakeLatency returns Latency based on passed time.Duration object.
func MakeLatency(d time.Duration, isGKE bool) Latency {
	if isGKE {
		return makeGKELatency(d)
	} else {
		return makeGAELatency(d)
	}
}

// makeGKELatency returns Latency struct for GKE based on passed time.Duration object.
func makeGKELatency(d time.Duration) Latency {
	return d.String()
}

// makeGAELatency returns Latency struct for Cloud Run and GAE based on passed time.Duration object.
func makeGAELatency(d time.Duration) Latency {
	nanos := d.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return GAELatency{
		Nanos:   int32(nanos),
		Seconds: secs,
	}
}

// remoteIP makes a best effort to compute the request client IP.
func remoteIP(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}

	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}

	return ip
}
