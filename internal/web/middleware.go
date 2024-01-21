package web

import (
	"github.com/gin-gonic/gin"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ResError(c, ErrMethodNotAllow)
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ResError(c, ErrNotFound)
	}
}
