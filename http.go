package zerodriver

import (
	"net/http"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/rs/zerolog"
)

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

type Latency struct {
	Nanos   int32 `json:"nanos"`
	Seconds int64 `json:"seconds"`
}

func (e *Event) HTTP(req *zapdriver.HTTPPayload) *zerolog.Event {
	return e.Event.Interface("httpRequest", req)
}

func NewHTTP(req *http.Request, res *http.Response) *zapdriver.HTTPPayload {
	return zapdriver.NewHTTP(req, res)
}

func MakeLatency(d time.Duration) Latency {
	nanos := d.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return Latency{
		Nanos:   int32(nanos),
		Seconds: secs,
	}
}
