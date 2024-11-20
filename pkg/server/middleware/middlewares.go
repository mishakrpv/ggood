package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authorize(handler func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		handler(c)
	})
}
