package zerodriver

const (
	unknownServiceName = "unknown"
	serviceContextKey  = "serviceContext"
)

// ServiceContext adds the correct service information adding the log line
// It is a required field if an error needs to be reported.
//
// see: https://cloud.google.com/error-reporting/reference/rest/v1beta1/ServiceContext
// see: https://cloud.google.com/error-reporting/docs/formatting-error-messages
func (e *Event) ServiceContext(serviceName string) *Event {
	e.Interface(serviceContextKey, map[string]string{"service": serviceName})
	return e
}
