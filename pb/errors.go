package pb

import "fmt"

// Errorf quickly creates a new error message for use with multi-errors.
func Errorf(code uint32, format string, a ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Errorf(format, a...).Error(),
	}
}

// Errore quickly creates a new error from another error.
func Errore(code uint32, err error) *Error {
	return &Error{Code: code, Message: err.Error()}
}

// Errors exposes a mechanism for quickly creating MultiError objects.
type Errors []*Error

// Serialize a list of errors into a multi-error.
func (e Errors) Serialize() *MultiError {
	return &MultiError{
		Errors: e,
	}
}

// Error ensures that pb.Error implements error
func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Error ensures that pb.MultiError implements error
func (e *MultiError) Error() string {
	return fmt.Sprintf("%d errors", len(e.Errors))
}
