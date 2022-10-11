package zerodriver

// operation is the complete payload that can be interpreted by Cloud Logging as
// an operation.
// see: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogEntryOperation
type operation struct {
	// Optional. An arbitrary operation identifier. Log entries with the same
	// identifier are assumed to be part of the same operation.
	ID string `json:"id"`

	// Optional. An arbitrary producer identifier. The combination of id and
	// producer must be globally unique. Examples for producer:
	// "MyDivision.MyBigCompany.com", "github.com/MyProject/MyApplication".
	Producer string `json:"producer"`

	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `json:"first"`

	// Optional. Set this to True if this is the last log entry in the operation.
	Last bool `json:"last"`
}

// Operation adds the correct Cloud Logging "operation" field.
//
// Additional information about a potentially long-running operation with which
// a log entry is associated.
func (e *Event) Operation(id, producer string, first, last bool) *Event {
	op := &operation{
		ID:       id,
		Producer: producer,
		First:    first,
		Last:     last,
	}

	e.Event.Interface("logging.googleapis.com/operation", op)
	return e
}

// OperationStart is a function for logging `Operation`. It should be called
// for the first operation log.
func (e *Event) OperationStart(id, producer string) *Event {
	return e.Operation(id, producer, true, false)
}

// OperationContinue is a function for logging `Operation`. It should be called
// for any non-start/end operation log.
func (e *Event) OperationContinue(id, producer string) *Event {
	return e.Operation(id, producer, false, false)
}

// OperationEnd is a function for logging `Operation`. It should be called
// for the last operation log.
func (e *Event) OperationEnd(id, producer string) *Event {
	return e.Operation(id, producer, false, true)
}
