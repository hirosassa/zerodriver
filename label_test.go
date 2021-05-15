package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	t.Parallel()

	l := Label("key", "value")
	assert.Equal(t, &label{key: "labels.key", value: "value"}, l)
}

func TestLabels(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().Labels(Label("foo", "bar"), Label("baz", "qux")).Msg("labels")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/labels", zerolog.Dict().
		Str("baz", "qux").
		Str("foo", "bar")).Msg("labels")
	expected := out.String()
	out.Reset()

	assert.Equal(t, expected, actual)
}

func TestLabelsMultiple(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	labeledLog := log.Info().Labels(Label("foo", "bar"))
	labeledLog.Labels(Label("baz", "qux")).Msg("labels")
	actual := out.String()
	out.Reset()

	log.Info().Dict("logging.googleapis.com/labels", zerolog.Dict().
		Str("baz", "qux").
		Str("foo", "bar")).Msg("labels")
	expected := out.String()
	out.Reset()

	assert.Equal(t, expected, actual)
}
