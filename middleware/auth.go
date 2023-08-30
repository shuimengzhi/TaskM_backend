package middleware

import (
	"github.com/gin-gonic/gin"
)

var UserId int

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
	}
}
