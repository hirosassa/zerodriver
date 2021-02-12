package zerodriver_test

import (
	"testing"

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
