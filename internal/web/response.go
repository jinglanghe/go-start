package web

import (
	"errors"
	"fmt"
)

type ResponseError struct {
	Code    int // error code
	Message string
	Status  int // http code
	Err     error
}

func (r *ResponseError) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	var v *ResponseError
	if errors.As(err, &v) {
		return v
	}
	return nil
}

func WrapResponse(err error, code, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		Err:     err,
		Status:  status,
	}
	return res
}

func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 400, msg, args...)
}

func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 500, msg, args...)
}

func NewResponse(code, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		Status:  status,
	}
	return res
}

func New400Response(msg string, args ...interface{}) error {
	return NewResponse(0, 400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewResponse(0, 500, msg, args...)
}

var (
	ErrInvalidToken    = NewResponse(9999, 401, "invalid signature")
	ErrNoPerm          = NewResponse(0, 401, "no permission")
	ErrNotFound        = NewResponse(0, 404, "not found")
	ErrMethodNotAllow  = NewResponse(0, 405, "method not allowed")
	ErrTooManyRequests = NewResponse(0, 429, "too many requests")
	ErrInternalServer  = NewResponse(0, 500, "internal server error")
	ErrBadRequest      = New400Response("bad request")
	ErrInvalidParent   = New400Response("not found parent node")
	ErrUserDisable     = New400Response("user forbidden")
)
