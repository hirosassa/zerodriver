package zerodriver

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	t.Parallel()

	req := &HTTPPayload{}
	var e Event
	event := e.HTTP(req)

	assert.Equal(t, e.Interface("httpreqest", req), event)
}

func TestNewHTTP(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
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
			&http.Request{ContentLength: 5, Body: ioutil.NopCloser(strings.NewReader("12345"))},
			nil,
			&HTTPPayload{RequestSize: "5"},
		},

		"ResponseSize": {
			nil,
			&http.Response{ContentLength: 5, Body: ioutil.NopCloser(strings.NewReader("12345"))},
			&HTTPPayload{ResponseSize: "5"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, NewHTTP(tt.req, tt.res)); diff != "" {
				t.Errorf("HTTPPayload differs (-got +want)\n%s", diff)
			}
		})
	}
}

func TestMakeLatency(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		d    time.Duration
		want Latency
	}{
		"zero": {
			time.Duration(0),
			Latency{Nanos: 0, Seconds: 0},
		},
		"10 sec": {
			time.Duration(10 * 1e9),
			Latency{Nanos: 0, Seconds: 10},
		},
		"sec and nano": {
			time.Duration(123 + 456*1e9),
			Latency{Nanos: 123, Seconds: 456},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, MakeLatency(tt.d))
		})
	}
}

func TestRemoteIP(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
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
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, remoteIP(tt.req))
		})
	}
}
