package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func Test_EventMsg(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().Msg("test")
	actual := out.String()
	out.Reset()

	logger.Info().Msg("test")
	expected := out.String()

	assert.Equal(t, expected, actual)
}
