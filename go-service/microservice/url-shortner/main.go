package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Model
type URL struct {
	OriginalURL string    `json:"url"`
	ShortCode   string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	Clicks      int       `json:"clicks"`
}

// In-memory store
var (
	urlStore = make(map[string]*URL)
	mu       sync.RWMutex
)

// Generate short code
func generateCode(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}

// Request
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// POST /shorten
func shortenHandler(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	code := generateCode(6)

	mu.Lock()
	urlStore[code] = &URL{
		OriginalURL: req.URL,
		ShortCode:   code,
		CreatedAt:   time.Now(),
		Clicks:      0,
	}
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

// GET /:code
func redirectHandler(c *gin.Context) {
	code := c.Param("code")

	mu.Lock()
	url, exists := urlStore[code]
	if exists {
		url.Clicks++
	}
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}

// GET /stats/:code
func statsHandler(c *gin.Context) {
	code := c.Param("code")

	mu.RLock()
	url, exists := urlStore[code]
	mu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, url)
}

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenHandler)
	r.GET("/stats/:code", statsHandler)
	r.GET("/:code", redirectHandler)

	r.Run(":8080")
}