package zerodriver

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strconv"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestWithServiceName(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger(WithServiceName("test-gcp-project"))
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Info().Msg("test")
	actual := out.String()
	out.Reset()

	log.Info().Dict("serviceContext", zerolog.Dict().
		Str("service", "test-gcp-project")).
		Msg("test")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestWithEventErrorReport(t *testing.T) {
	frame := runtime.Frame{
		PC:   0,
		File: "/a/b/c/dummy.go",
		Line: 10,
	}
	mockCaller(t, frame, true)

	tests := []struct {
		name   string
		logger *Logger
		level  zerolog.Level
		want   map[string]interface{}
	}{
		{
			name: "Add fields when log level is error",
			logger: NewProductionLogger(
				WithServiceName("test-project"),
				WithReportAllErrors(),
			),
			level: zerolog.ErrorLevel,
			want: map[string]interface{}{
				"serviceContext": map[string]string{
					"service": "test-project",
				},
				"context": map[string]interface{}{
					"reportLocation": map[string]interface{}{
						"filePath":     frame.File,
						"lineNumber":   strconv.Itoa(frame.Line),
						"functionName": "",
					},
				},
			},
		},
		{
			name:   "Add fields with unknown service name when service name is not set",
			logger: NewProductionLogger(WithReportAllErrors()),
			level:  zerolog.ErrorLevel,
			want: map[string]interface{}{
				"serviceContext": map[string]string{
					"service": "unknown",
				},
				"context": map[string]interface{}{
					"reportLocation": map[string]interface{}{
						"filePath":     frame.File,
						"lineNumber":   strconv.Itoa(frame.Line),
						"functionName": "",
					},
				},
			},
		},
		{
			name:   "Don't add fields when log level is below error",
			logger: NewProductionLogger(WithReportAllErrors()),
			level:  zerolog.InfoLevel,
			want:   map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			log := tt.logger
			logger := zerolog.New(out).With().Logger()
			log.Logger = &logger

			log.WithLevel(tt.level).Msg("test")
			actual := make(map[string]interface{})
			err := json.Unmarshal([]byte(out.String()), &actual)
			assert.NoError(t, err)
			out.Reset()

			log.WithLevel(tt.level).Fields(tt.want).Msg("test")
			expected := make(map[string]interface{})
			err = json.Unmarshal([]byte(out.String()), &expected)
			assert.NoError(t, err)

			assert.Equal(t, expected, actual)
		})
	}
}
