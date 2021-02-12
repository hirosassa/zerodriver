package zerodriver_test

import (
	"errors"
	"testing"

	"github.com/hirosassa/zerodriver"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	logger := zerodriver.New(false)
	assert.IsType(t, &zerodriver.Logger{}, logger)
}

func TestLoggers(t *testing.T) {
	t.Parallel()

	log := zerodriver.New(false)

	var tests = map[string]struct {
		res  *zerodriver.Event
		want *zerodriver.Event
	}{
		"trace": {
			res:  log.Trace(),
			want: &zerodriver.Event{log.Logger.Trace()},
		},
		"debug": {
			res:  log.Debug(),
			want: &zerodriver.Event{log.Logger.Debug()},
		},
		"info": {
			res:  log.Info(),
			want: &zerodriver.Event{log.Logger.Info()},
		},
		"warn": {
			res:  log.Warn(),
			want: &zerodriver.Event{log.Logger.Warn()},
		},
		"error": {
			res:  log.Error(),
			want: &zerodriver.Event{log.Logger.Error()},
		},
		"err": {
			res:  log.Err(errors.New("some error")),
			want: &zerodriver.Event{log.Logger.Err(errors.New("some error"))},
		},
		"log": {
			res:  log.Log(),
			want: &zerodriver.Event{log.Logger.Log()},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.res)
		})
	}
}
