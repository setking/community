package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, Authorization, X-CSRF-Token, AccessToken")
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
