package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().Operation("id", "producer", true, false).Msg("operation")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/operation", zerolog.Dict().
		Str("id", "id").
		Str("producer", "producer").
		Bool("first", true).
		Bool("last", false)).
		Msg("operation")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestOperationStart(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().OperationStart("id", "producer").Msg("operation start")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/operation", zerolog.Dict().
		Str("id", "id").
		Str("producer", "producer").
		Bool("first", true).
		Bool("last", false)).Msg("operation start")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestOperationContinue(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().OperationContinue("id", "producer").Msg("operation continue")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/operation", zerolog.Dict().
		Str("id", "id").
		Str("producer", "producer").
		Bool("first", false).
		Bool("last", false)).Msg("operation continue")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestOperationEnd(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().OperationEnd("id", "producer").Msg("operation end")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/operation", zerolog.Dict().
		Str("id", "id").
		Str("producer", "producer").
		Bool("first", false).
		Bool("last", true)).Msg("operation end")
	expected := out.String()

	assert.Equal(t, expected, actual)
}
