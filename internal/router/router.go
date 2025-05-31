package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ronihdz/simple-gin-lambda/internal/middleware"
)

func New(basePath string) *gin.Engine {
	r := gin.New()

	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.Use(gin.Recovery(), middleware.RequestLogger())

	root := r.Group(basePath)

	// ───────────────  Endpoints  ───────────────
	root.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "📚 BooksGo API",
			"hint":    basePath + "/health",
		})
	})

	root.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	root.POST("/echo", func(c *gin.Context) {
		var body map[string]any
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, body)
	})

	root.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{"serverTime": time.Now().Format(time.RFC3339)})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":       "ruta no encontrada",
			"requestPath": c.Request.URL.Path,
		})
	})

	return r
}
