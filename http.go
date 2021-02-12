package zerodriver

import (
	"net/http"

	"github.com/blendle/zapdriver"
	"github.com/rs/zerolog"
)

func (e *Event) HTTP(req *zapdriver.HTTPPayload) *zerolog.Event {
	ze := zerolog.Event(*e.Event)
	return ze.Interface("httpRequest", req)
}

func NewHTTP(req *http.Request, res *http.Response) *zapdriver.HTTPPayload {
	return zapdriver.NewHTTP(req, res)
}
