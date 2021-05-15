package zerodriver

import (
	"fmt"

	"github.com/rs/zerolog"
)

// TraceContext adds the "trace", "span", "trace_sampled" fields to the *zerolog.Event context.
//
// see: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
func (e *Event) TraceContext(trace string, spanId string, sampled bool, projectID string) *zerolog.Event {
	return e.event.
		Str("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", projectID, trace)).
		Str("logging.googleapis.com/spanId", spanId).
		Bool("logging.googleapis.com/trace_sampled", sampled)
}
