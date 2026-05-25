package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// disable redirect issues (important for gateway setups)
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	// match /users and anything after it
	r.Any("/users/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "" {
			path = "/"
		}

		c.String(200, fmt.Sprintf("User Service ✅ Path: %s", path))
	})

	fmt.Println("User service running on :9001")
	r.Run(":9001")
}