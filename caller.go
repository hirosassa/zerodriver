package zerodriver

import (
	"runtime"
	"testing"
)

var caller = realCaller

func realCaller(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}

func mockCaller(t *testing.T, frame runtime.Frame, ok bool) {
	t.Helper()

	caller = func(skip int) (uintptr, string, int, bool) {
		return frame.PC, frame.File, frame.Line, ok
	}
}
