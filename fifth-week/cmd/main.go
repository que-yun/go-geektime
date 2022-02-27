package main

import (
	"github.com/gin-gonic/gin"
	"go-geektime/fifth-week/pkg/middleware"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Wrapper(100, 10, 0.5, time.Second))
	engine.Handle("GET", "/", func(c *gin.Context) {
		c.Writer.WriteString("helloWorld")
	})

	engine.Handle("GET", "/err", func(c *gin.Context) {
		c.Writer.WriteHeader(404)
	})

	engine.Run(":8080")
}
