package main

import (
	"database/sql"
	"error-demo/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				fmt.Printf("ErrNoRows: %+v\n", err)
				//c.Status(http.StatusInternalServerError)
				c.AbortWithError(http.StatusInternalServerError, errors.Unwrap(err))
				return
			}
			fmt.Printf("%+v\n", err)
		}
	}
}

func main() {
	engine := gin.New()
	engine.Use(LoggerMiddleware)
	engine.GET("/testError", wrapper(handler.FindUser))
	if err := engine.Run(":8080"); err != nil {
		panic("[error-demo] gin.run fatal")
	}
}

func LoggerMiddleware(c *gin.Context) {
	c.Next()
}
