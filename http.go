package zerodriver

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

// HTTPPayload is the struct consists of http request related components.
// Details are in following link.
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
type HTTPPayload struct {
	RequestMethod                  string  `json:"requestMethod,omitempty"`
	RequestURL                     string  `json:"requestUrl,omitempty"`
	RequestSize                    string  `json:"requestSize,omitempty"`
	Status                         int     `json:"status,omitempty"`
	ResponseSize                   string  `json:"responseSize,omitempty"`
	UserAgent                      string  `json:"userAgent,omitempty"`
	RemoteIP                       string  `json:"remoteIp,omitempty"`
	ServerIP                       string  `json:"serverIp,omitempty"`
	Referer                        string  `json:"referer,omitempty"`
	Latency                        Latency `json:"latency,omitempty"`
	CacheLookup                    bool    `json:"cacheLookup,omitempty"`
	CacheHit                       bool    `json:"cacheHit,omitempty"`
	CacheValidatedWithOriginServer bool    `json:"cacheValidatedWithOriginServer,omitempty"`
	CacheFillBytes                 string  `json:"cacheFillBytes,omitempty"`
	Protocol                       string  `json:"protocol,omitempty"`
}

// The request processing latency on the server, from the time the request was
// received until the response was sent.
type Latency struct {
	Nanos   int32 `json:"nanos"`
	Seconds int64 `json:"seconds"`
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
