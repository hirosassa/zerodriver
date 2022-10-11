package zerodriver

import (
	"runtime"
	"strconv"
)

const contextKey = "context"

type reportLocation struct {
	File     string `json:"filePath"`
	Line     string `json:"lineNumber"`
	Function string `json:"functionName"`
}

// reportContext is the context information attached to a log for reporting errors
type reportContext struct {
	ReportLocation reportLocation `json:"reportLocation"`
}

func newReportContext(pc uintptr, file string, line int) *reportContext {
	var function string
	if fn := runtime.FuncForPC(pc); fn != nil {
		function = fn.Name()
	}

	context := &reportContext{
		ReportLocation: reportLocation{
			File:     file,
			Line:     strconv.Itoa(line),
			Function: function,
		},
	}

	return context
}

// ErrorReport adds the correct Stackdriver "context" field for getting the log line
// reported as error.
//
// see: https://cloud.google.com/error-reporting/docs/formatting-error-messages
func (e *Event) ErrorReport(pc uintptr, file string, line int, ok bool) *Event {
	if !ok {
		return e
	}

	e.Interface(contextKey, newReportContext(pc, file, line))
	return e
}
