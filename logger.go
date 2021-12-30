package zerodriver

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type config struct {
	serviceName     string
	reportAllErrors bool
}

type Logger struct {
	*zerolog.Logger
	config *config
}

type Event struct {
	*zerolog.Event

	config *config
	level  zerolog.Level
}

// See: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
var logLevelSeverity = map[zerolog.Level]string{
	zerolog.DebugLevel: "DEBUG",
	zerolog.InfoLevel:  "INFO",
	zerolog.WarnLevel:  "WARNING",
	zerolog.ErrorLevel: "ERROR",
	zerolog.PanicLevel: "CRITICAL",
	zerolog.FatalLevel: "CRITICAL",
}

// NewProductionLogger returns a configured logger for production.
// It outputs info level and above logs with sampling.
func NewProductionLogger(opts ...Option) *Logger {
	logLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity[l]
	}
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// default sampler
	sampler := &zerolog.BasicSampler{N: 1}

	logger := zerolog.New(os.Stderr).Sample(sampler).With().Timestamp().Logger()
	logger.With().Caller().Caller()
	return &Logger{Logger: &logger, config: newConfig(opts...)}
}

// NewDevelopmentLogger returns a configured logger for development.
// It outputs debug level and above logs, and sampling is disabled.
func NewDevelopmentLogger(opts ...Option) *Logger {
	logLevel := zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity[l]
	}
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{Logger: &logger, config: newConfig(opts...)}
}

// To use method chain we need followings

func (l *Logger) Trace() *Event {
	e := l.Logger.Trace()
	return &Event{Event: e, config: l.config, level: zerolog.TraceLevel}
}

func (l *Logger) Debug() *Event {
	e := l.Logger.Debug()
	return &Event{Event: e, config: l.config, level: zerolog.DebugLevel}
}

func (l *Logger) Info() *Event {
	e := l.Logger.Info()
	return &Event{Event: e, config: l.config, level: zerolog.InfoLevel}
}

func (l *Logger) Warn() *Event {
	e := l.Logger.Warn()
	return &Event{Event: e, config: l.config, level: zerolog.WarnLevel}
}

func (l *Logger) Error() *Event {
	e := l.Logger.Error()
	return &Event{Event: e, config: l.config, level: zerolog.ErrorLevel}
}

func (l *Logger) Err(err error) *Event {
	e := l.Logger.Error().Err(err)
	return &Event{Event: e, config: l.config, level: zerolog.ErrorLevel}
}

func (l *Logger) Fatal() *Event {
	e := l.Logger.Fatal()
	return &Event{Event: e, config: l.config, level: zerolog.FatalLevel}
}

func (l *Logger) Panic() *Event {
	e := l.Logger.Panic()
	return &Event{Event: e, config: l.config, level: zerolog.PanicLevel}
}

func (l *Logger) WithLevel(level zerolog.Level) *Event {
	e := l.Logger.WithLevel(level)
	return &Event{Event: e, config: l.config, level: level}
}

func (l *Logger) Log() *Event {
	e := l.Logger.Log()
	return &Event{Event: e, config: l.config, level: zerolog.NoLevel}
}

func (l *Logger) Print(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

func (l Logger) Write(p []byte) (n int, err error) {
	n, err = l.Logger.Write(p)
	return n, err
}
