package zerodriver

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewProduction(t *testing.T) {
	logger := NewProductionLogger()
	assert.IsType(t, &Logger{}, logger)
}

func TestNewDevelopment(t *testing.T) {
	logger := NewDevelopmentLogger()
	assert.IsType(t, &Logger{}, logger)
}

func TestLoggers(t *testing.T) {
	t.Parallel()

	log := NewProductionLogger()

	var tests = map[string]struct {
		res     *Event
		want    *Event
		wantErr error
	}{
		"trace": {
			res:  log.Trace(),
			want: &Event{log.Logger.Trace()},
		},
		"debug": {
			res:  log.Debug(),
			want: &Event{log.Logger.Debug()},
		},
		"info": {
			res:  log.Info(),
			want: &Event{log.Logger.Info()},
		},
		"warn": {
			res:  log.Warn(),
			want: &Event{log.Logger.Warn()},
		},
		"error": {
			res:  log.Error(),
			want: &Event{log.Logger.Error()},
		},
		"err": {
			res:  log.Err(errors.New("some error")),
			want: &Event{log.Logger.Err(errors.New("some error"))},
		},
		"with level": {
			res:  log.WithLevel(zerolog.InfoLevel),
			want: &Event{log.Logger.WithLevel(zerolog.InfoLevel)},
		},
		"log": {
			res:  log.Log(),
			want: &Event{log.Logger.Log()},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.res)
		})
	}
}

func TestFatal(t *testing.T) {
	t.Parallel()

	// patch os.Exit
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	defer func() {
		_ = recover()
		actual := out.String()
		out.Reset()

		log.WithLevel(zerolog.FatalLevel).Msg("fatal")
		expected := out.String()
		assert.Equal(t, expected, actual)
	}()

	log.Fatal().Msg("fatal")
}

func TestPanic(t *testing.T) {
	t.Parallel()

	// replace writer
	log := NewProductionLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	defer func() {
		_ = recover()
		actual := out.String()
		out.Reset()

		log.WithLevel(zerolog.FatalLevel).Msg("panic")
		expected := out.String()
		assert.Equal(t, expected, actual)
	}()

	log.Panic().Msg("panic")
}

func TestPrint(t *testing.T) {
	// replace writer
	log := NewDevelopmentLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Print("print")
	actual := out.String()
	out.Reset()

	log.Debug().Msg("print")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestPrintf(t *testing.T) {
	// replace writer
	log := NewDevelopmentLogger()
	out := &bytes.Buffer{}
	logger := zerolog.New(out).With().Logger()
	log.Logger = &logger

	log.Printf("print: %s", "hello")
	actual := out.String()
	out.Reset()

	log.Debug().Msgf("print: %s", "hello")
	expected := out.String()

	assert.Equal(t, expected, actual)
}

func TestWrite(t *testing.T) {
	t.Parallel()

	log := NewProductionLogger()
	n, err := log.Write([]byte("abc"))
	assert.Equal(t, 3, n)
	assert.NoError(t, err)
}
