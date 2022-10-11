package zerodriver

import "github.com/rs/zerolog"

func (e *Event) Msg(msg string) {
	if e.config.serviceName != "" {
		e.ServiceContext(e.config.serviceName)
	}
	if e.config.reportAllErrors && e.level >= zerolog.ErrorLevel {
		if e.config.serviceName == "" {
			// A service name was not set but error report needs it
			// So attempt to add a generic service name
			e.ServiceContext(unknownServiceName)
		}
		e.ErrorReport(caller(zerolog.CallerSkipFrameCount))
	}
	e.Event.Msg(msg)
}
