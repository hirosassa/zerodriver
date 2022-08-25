package zerodriver

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	req := &HTTPPayload{RequestURL: "https://example.com"}
	log.Info().HTTP(req).Msg("http request")
	actual := out.String()
	out.Reset()

	log.Info().Dict("httpRequest", zerolog.Dict().
		Str("requestMethod", "").
		Str("requestUrl", "https://example.com").
		Str("requestSize", "").
		Int("status", 0).
		Str("responseSize", "").
		Str("userAgent", "").
		Str("remoteIp", "").
		Str("serverIp", "").
		Str("referer", "").
		Interface("latency", nil).
		Bool("cacheLookup", false).
		Bool("cacheHit", false).
		Bool("cacheValidatedWithOriginServer", false).
		Str("cacheFillBytes", "").
		Str("protocol", "")).
		Msg("http request")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestNewHTTP(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		req  *http.Request
		res  *http.Response
		want *HTTPPayload
	}{
		"empty": {
			nil,
			nil,
			&HTTPPayload{},
		},
		"RequestURL": {
			&http.Request{URL: &url.URL{Host: "example.com", Scheme: "https"}},
			nil,
			&HTTPPayload{RequestURL: "https://example.com"},
		},
		"RequestSize": {
			&http.Request{ContentLength: 5, Body: io.NopCloser(strings.NewReader("12345"))},
			nil,
			&HTTPPayload{RequestSize: "5"},
		},

		"ResponseSize": {
			nil,
			&http.Response{ContentLength: 5, Body: io.NopCloser(strings.NewReader("12345"))},
			&HTTPPayload{ResponseSize: "5"},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(tt.want, NewHTTP(tt.req, tt.res)); diff != "" {
				t.Errorf("HTTPPayload differs (-got +want)\n%s", diff)
			}
		})
	}
}

func TestMakeLatency(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		d     time.Duration
		isGKE bool
		want  Latency
	}{
		"gke": {
			time.Duration(0),
			true,
			"0s",
		},
		"gae": {
			time.Duration(0),
			false,
			GAELatency{Seconds: 0, Nanos: 0},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, MakeLatency(tt.d, tt.isGKE))
		})
	}
}

func TestMakeGAELatency(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		d    time.Duration
		want GAELatency
	}{
		"zero": {
			time.Duration(0),
			GAELatency{Nanos: 0, Seconds: 0},
		},
		"10 sec": {
			time.Duration(10 * 1e9),
			GAELatency{Nanos: 0, Seconds: 10},
		},
		"sec and nano": {
			time.Duration(123 + 456*1e9),
			GAELatency{Nanos: 123, Seconds: 456},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, makeGAELatency(tt.d))
		})
	}
}

func TestRemoteIP(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		req  *http.Request
		want string
	}{
		"empty": {
			&http.Request{},
			"",
		},
		"x-forwarded-for header": {
			&http.Request{Header: map[string][]string{
				"X-Forwarded-For": {"0.0.0.0"},
			}},
			"0.0.0.0",
		},
		"remote address": {
			&http.Request{RemoteAddr: "0.0.0.0"},
			"0.0.0.0",
		},
		"remote address with port": {
			&http.Request{RemoteAddr: "0.0.0.0:3000"},
			"0.0.0.0",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, remoteIP(tt.req))
		})
	}
}
