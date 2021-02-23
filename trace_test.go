package zerodriver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraceContext(t *testing.T) {
	t.Parallel()

	var e Event
	event := e.TraceContext("06796866738c859f2f19b7cfb3214824", "000000000000004a", false, "my-projectid")
	assert.Equal(t,
		e.Str("logging.googleapis.com/trace", "projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824").
			Str("logging.googleapis.com/spanId", "000000000000004a").
			Bool("logging.googleapis.com/trace_sampled", false),
		event)
}
