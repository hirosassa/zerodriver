package zerodriver

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

type Event struct {
	*zerolog.Event
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

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity[l]
	}
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{&logger}
}


// To use method chain we need followings

func (l *Logger) Trace() *Event {
	e := l.Logger.Trace()
	return &Event{e}
}

func (l *Logger) Debug() *Event {
	e := l.Logger.Debug()
	return &Event{e}
}

func (l *Logger) Info() *Event {
	e := l.Logger.Info()
	return &Event{e}
}

func (l *Logger) Warn() *Event {
	e := l.Logger.Warn()
	return &Event{e}
}

func (l *Logger) Error() *Event {
	e := l.Logger.Error()
	return &Event{e}
}

func (l *Logger) Err(err error) *Event {
	e := l.Logger.Error().Err(err)
	return &Event{e}
}

func (l *Logger) Fatal() *Event {
	e := l.Logger.Fatal()
	return &Event{e}
}

func (l *Logger) Panic() *Event {
	e := l.Logger.Panic()
	return &Event{e}
}

func (l *Logger) WithLevel(level zerolog.Level) *Event {
	e := l.Logger.WithLevel(level)
	return &Event{e}
}

func (l *Logger) Log() *Event {
	e := l.Logger.Log()
	return &Event{e}
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
