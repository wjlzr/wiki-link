package middleware

import (
	"github.com/gin-gonic/gin"
)

// 自定义参数拦截
func CustomIntercept() gin.HandlerFunc {

	return func(c *gin.Context) {
		xclient := c.Request.Header.Get("x-client")
		if xclient == "" {
			c.Request.Header.Set("x-client", "app")
		}
		c.Next()
	}
}
