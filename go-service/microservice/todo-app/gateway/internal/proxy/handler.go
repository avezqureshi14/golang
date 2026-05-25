package gateway

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func Handler(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {

		// forward user info
		if userID, exists := c.Get("userID"); exists {
			c.Request.Header.Set("X-User-ID", userID.(string))
		}

		// strip prefix
		c.Request.URL.Path = c.Param("path")

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}