package pb

import "fmt"

//===========================================================================
// Create Single Errors
//===========================================================================

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

// Error ensures that pb.Error implements error
func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

//===========================================================================
// Multiple Errors
//===========================================================================

// Errors exposes a mechanism for quickly creating MultiError objects.
type Errors []*Error

// Serialize a list of errors into a multi-error.
func (e Errors) Serialize() *MultiError {
	return &MultiError{
		Errors: e,
	}
}

// Error ensures that pb.Errors implements error
func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	return fmt.Sprintf("%d errors", len(e))
}

//===========================================================================
// MultiError functionality
//===========================================================================

// Error ensures that pb.MultiError implements error
func (e *MultiError) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	return fmt.Sprintf("%d errors", len(e.Errors))
}

// Deserialize returns the errors object and functionality
func (e *MultiError) Deserialize() Errors {
	errs := make(Errors, 0, len(e.Errors))
	for _, e := range e.Errors {
		errs = append(errs, e)
	}
	return errs
}
