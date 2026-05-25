package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// create proxy
func newProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	r := gin.Default()

	// services
	userService := newProxy("http://localhost:9001")
	todoService := newProxy("http://localhost:9002")

	// routes
	r.Any("/users/*path", func(c *gin.Context) {
		userService.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/todos/*path", func(c *gin.Context) {
		todoService.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8080")
}