package errors

import "net/http"

func (err *Error) HttpStatus() int {
	switch err.code {
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeInvalid:
		return http.StatusBadRequest
	default:
		return 500
	}
}
