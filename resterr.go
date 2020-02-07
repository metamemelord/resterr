/*
Package resterr makes it easy to report errors between the layers of your RESTful service without putting multiple checks
for deciding http status codes.
*/
package resterr

import "fmt"

var showStatusCodeInError bool

// RESTErr interface gives error and status code
type RESTErr interface {
	error
	StatusCode() uint

	WithStatusCode(uint) RESTErr
}

// New returns a new RESTErr
func New(errorString string) RESTErr {
	return &restErr{errorString: errorString, statusCode: 500}
}

// Errorf returns a RESTErr with status code 500, format, and arguments
func Errorf(format string, args ...interface{}) RESTErr {
	return &restErr{errorString: fmt.Sprintf(format, args...), statusCode: 500}
}

// ErrorfWithStatusCode returns a RESTErr with status code, format, and arguments
func ErrorfWithStatusCode(statusCode uint, format string, args ...interface{}) RESTErr {
	return &restErr{errorString: fmt.Sprintf(format, args...), statusCode: statusCode}
}

// ShowStatusCodeInError enables showing status code package wide.
// Status code is shown in error message only if it's set
func ShowStatusCodeInError(value bool) {
	showStatusCodeInError = value
}

type restErr struct {
	statusCode  uint
	errorString string
}

func (re *restErr) Error() string {
	if showStatusCodeInError && re.statusCode != 0 {
		return fmt.Sprintf("%d: %s", re.statusCode, re.errorString)
	}
	return re.errorString
}

func (re *restErr) StatusCode() uint {
	return re.statusCode
}

func (re *restErr) WithStatusCode(statusCode uint) RESTErr {
	if statusCode < 100 || statusCode >= 600 {
		panic("Invalid status code")
	}
	re.statusCode = statusCode
	return re
}
