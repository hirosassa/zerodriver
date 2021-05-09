package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestTraceContext(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().TraceContext("06796866738c859f2f19b7cfb3214824", "000000000000004a", false, "my-projectid").Msg("trace")
	actual := string(out.Bytes())
	out.Reset()

	log.Info().Str("logging.googleapis.com/trace", "projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824").
		Str("logging.googleapis.com/spanId", "000000000000004a").
		Bool("logging.googleapis.com/trace_sampled", false).
		Msg("trace")
	expected := string(out.Bytes())
	assert.Equal(t, expected, actual)
}
