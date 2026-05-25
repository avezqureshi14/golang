package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 🔥 important for gateway setups (avoid redirect loops)
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	// match /todos and anything after it
	r.Any("/todos/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "" {
			path = "/"
		}

		c.String(200, fmt.Sprintf("Todo Service ✅ Path: %s", path))
	})

	fmt.Println("Todo service running on :9002")
	r.Run(":9002")
}