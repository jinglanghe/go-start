package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinglanghe/go-start/internal/errorx"
	"github.com/jinglanghe/go-start/utils/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	prefix     = "go-start"
	ReqBodyKey = prefix + "/req-body"
	ResBodyKey = prefix + "/res-body"
)

type APISuccessResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ParseParamID returns the value of the URL param
func ParseParamID(c *gin.Context, key string) int {
	val := c.Param(key)
	id, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0
	}
	return (int)(id)
}

func ParseHeader(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindHeader(obj); err != nil {
		return Wrap400Response(err, fmt.Sprintf("Parse request header failed: %s", err.Error()))
	}
	return nil
}

// ParseJSON data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return Wrap400Response(err, fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseQuery parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return Wrap400Response(err, fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

// ParseForm data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return Wrap400Response(err, fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

func ParseUri(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		return Wrap400Response(err, fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

// ResSuccess data object
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON data with status code
func ResJSON(c *gin.Context, status int, v interface{}) {
	res := APISuccessResp{
		Code: errorx.ActionSuccess.Code(),
		Msg:  "success",
		Data: v,
	}
	c.JSON(http.StatusOK, res)

}

type APIException struct {
	StatusCode int               `json:"-"`
	Code       int               `json:"code"`
	Msg        string            `json:"msg"`
	Data       map[string]string `json:"data"`
}

// ResError object and parse error status code
func ResError(c *gin.Context, err error, status ...int) {
	apiException := handleErr(err)
	c.JSON(apiException.StatusCode, apiException)
}

func handleErr(err error) *APIException {
	log.Error(err).Send()
	var apiException *APIException
	var baseError errorx.BaseError
	switch {
	case errors.As(err, &baseError):
		apiException = convertToApiException(err.(errorx.BaseError))
	}
	return apiException
}

func convertToApiException(err errorx.BaseError) *APIException {
	return &APIException{
		StatusCode: err.ErrorCode().StatusCode(),
		Code:       err.ErrorCode().Code(),
		Msg:        err.Error(),
		Data:       err.ErrorData(),
	}
}
