package zerodriver

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestErrorReport(t *testing.T) {
	t.Parallel()

	t.Run("Add error reporting context fields when caller is ok", func(t *testing.T) {
		// replace writer
		log := NewProductionLogger()
		out := &bytes.Buffer{}
		logger := zerolog.New(out).With().Logger()
		log.Logger = &logger

		log.Info().ErrorReport(0, "a/b/c/dummy.go", 10, true).Msg("test")
		actual := out.String()
		out.Reset()

		log.Info().Dict("context", zerolog.Dict().Dict(
			"reportLocation", zerolog.Dict().
				Str("filePath", "a/b/c/dummy.go").
				Str("lineNumber", "10").
				Str("functionName", ""),
		)).
			Msg("test")
		expected := out.String()

		assert.Equal(t, expected, actual)
	})

	t.Run("Add error reporting context fields when failed to get caller", func(t *testing.T) {
		// replace writer
		log := NewProductionLogger()
		out := &bytes.Buffer{}
		logger := zerolog.New(out).With().Logger()
		log.Logger = &logger

		log.Info().ErrorReport(0, "", 0, false).Msg("test")
		actual := out.String()
		out.Reset()

		log.Info().Msg("test")
		expected := out.String()

		assert.Equal(t, expected, actual)
	})
}
