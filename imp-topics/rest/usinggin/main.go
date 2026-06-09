package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users  = []User{}
	mu     sync.Mutex
	nextID = 1
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var u User

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		u.ID = nextID
		nextID++
		users = append(users, u)
		mu.Unlock()

		c.JSON(http.StatusOK, u)
	})

	r.Run(":8080")
}
