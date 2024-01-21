package errorx

import "net/http"

const (
	ActionSuccess ErrorType = 0

	NotFoundError     ErrorType = 320810001
	UnknownError      ErrorType = 320810002
	ServerError       ErrorType = 320810003
	InvalidParamError ErrorType = 320810004

	DBNotFoundError  ErrorType = 320820001
	DBOperationError ErrorType = 320820000

	DefaultAppErr ErrorType = 320830001

	ForbiddenAppErr ErrorType = 320840001

	NOPStatusAppErr ErrorType = 320850001
)

type BaseError interface {
	error
	ErrorCode() ErrorType
	ErrorData() map[string]string
}

type ErrorType int

func (e ErrorType) Code() int {
	return int(e)
}

func (e ErrorType) StatusCode() int {
	if statusCode, ok := error2StatusCodeMap[e]; ok {
		return statusCode
	}
	return http.StatusBadRequest
}

var error2StatusCodeMap = map[ErrorType]int{
	InvalidParamError: http.StatusBadRequest,
	DBNotFoundError:   http.StatusBadRequest,
	DBOperationError:  http.StatusBadRequest,
	DefaultAppErr:     http.StatusBadRequest,
	ForbiddenAppErr:   http.StatusBadRequest,
	NOPStatusAppErr:   http.StatusBadRequest,
	ServerError:       http.StatusInternalServerError,
	NotFoundError:     http.StatusNotFound,
}
