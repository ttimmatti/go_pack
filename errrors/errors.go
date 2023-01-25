package errrors

import "fmt"

func WrapErrorF(orig error, code ErrorCode, cause string, a ...interface{}) error {
	return &Error{
		Orig: orig,
		code: code,
		msg:  fmt.Sprintf(cause, a...),
	}
}

// Error returns the message, when wrapping errors the wrapped error is returned.
func (e *Error) Error() string {
	if e.Orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.Orig)
	}

	return e.msg
}

// NewErrorf instantiates a new error.
func NewErrorf(code ErrorCode, cause string, a ...interface{}) error {
	return WrapErrorF(nil, code, cause, a...)
}

// Code returns the code representing this error.
func (e *Error) Code() ErrorCode {
	return e.code
}

// can be altered
type Error struct {
	Orig error
	msg  string
	code ErrorCode
}

// which errors are there
type ErrorCode uint

////////////////////////////////////////
// to read error as Error

// var ierr *Error

// 	errors.As(err, &ierr)

// 	log.Println(ierr.code)

//

// INIT ERROR CODE
// const (
// 	ErrorCodeUnknown ErrorCode = iota
// 	ErrorCodeFailure
// 	ErrorCodeNotFound
// 	ErrorCodeInvalidArgument
// )
