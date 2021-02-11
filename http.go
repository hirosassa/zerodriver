package zerodriver

import (
	"net/http"

	"github.com/blendle/zapdriver"
	"github.com/rs/zerolog"
)

func HTTP(req *zapdriver.HTTPPayload) *zerolog.Event {
	return zerolog.Dict().Interface("httpRequest", req)
}

func NewHTTP(req *http.Request, res *http.Response) *zapdriver.HTTPPayload {
	return zapdriver.NewHTTP(req, res)
}
