package zerodriver_test

import (
	"testing"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/hirosassa/zerodriver"
	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	t.Parallel()

	req := &zapdriver.HTTPPayload{}
	var e zerodriver.Event
	event := e.HTTP(req)

	assert.Equal(t, e.Interface("httpreqest", req), event)
}

func TestMakeLatency(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		d    time.Duration
		want zerodriver.Latency
	}{
		"zero": {
			time.Duration(0),
			zerodriver.Latency{Nanos: 0, Seconds: 0},
		},
		"10 sec": {
			time.Duration(10 * 1e9),
			zerodriver.Latency{Nanos: 0, Seconds: 10},
		},
		"sec and nano": {
			time.Duration(123 + 456*1e9),
			zerodriver.Latency{Nanos: 123, Seconds: 456},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, zerodriver.MakeLatency(tt.d))
		})
	}
}
