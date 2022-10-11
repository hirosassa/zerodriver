package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestServiceContext(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().ServiceContext("test-gcp-project").Msg("test")
	actual := out.String()
	out.Reset()

	log.Info().Dict("serviceContext", zerolog.Dict().
		Str("service", "test-gcp-project")).
		Msg("test")
	expected := out.String()

	assert.Equal(t, expected, actual)
}
