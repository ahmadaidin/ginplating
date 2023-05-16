package errors

import (
	"errors"
	"fmt"
)

type ErrCode string

const (
	ErrCodeNotFound ErrCode = "not_found"
	ErrCodeInternal ErrCode = "internal"
	ErrCodeConflict ErrCode = "conflict"
	ErrCodeInvalid  ErrCode = "invalid"
	ErrUnknown      ErrCode = ""
)

const separator = ". "

type Error struct {
	// code will be used to identify the error type
	code ErrCode

	// err will be used to identify the error message
	err error

	// messaege will be used to identify the error message
	message string
}

func (e *Error) PrependMsg(message string) *Error {
	e.message = fmt.Sprintf("%s%s%s", message, separator, e.message)
	return e
}

func (e *Error) AppendMsg(message string) *Error {
	e.message = fmt.Sprintf("%s%s%s", e.message, separator, message)
	return e
}

func (e Error) Err() error {
	return e.err
}

func (e Error) Valid() bool {
	return e.err != nil
}

func (e Error) Code() ErrCode {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e *Error) PrependError(err error) *Error {
	e.err = errors.Join(err, e.err)
	return e
}

func (e *Error) AppendError(err error) *Error {
	e.err = errors.Join(e.err, err)
	return e
}

func (e *Error) SetCode(code ErrCode) *Error {
	e.code = code
	return e
}

type joinError interface {
	Unwrap() []error
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, err: %v message: %v", e.code, e.err, e.message)
}

func (e *Error) UnwrapErr() []error {
	if e.err == nil {
		return nil
	}
	if je, ok := e.err.(joinError); ok {
		return je.Unwrap()
	}
	return []error{e.err}
}

func NewInternalError(err error, message string) Error {
	return Error{
		code:    ErrCodeInternal,
		err:     err,
		message: message,
	}
}

func NewNotFoundError(err error, message string) Error {
	return Error{
		code:    ErrCodeNotFound,
		err:     err,
		message: message,
	}
}

func NewConflictError(err error, message string) Error {
	return Error{
		code:    ErrCodeConflict,
		err:     err,
		message: message,
	}
}

func NewInvalidError(err error, message string) Error {
	return Error{
		code:    ErrCodeInvalid,
		err:     err,
		message: message,
	}
}

func NewUnknownError(err error, message string) Error {
	return Error{
		code:    ErrUnknown,
		err:     err,
		message: message,
	}
}

func New(text string) error {
	return errors.New(text)
}
