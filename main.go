package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	webserver := gin.Default()

	webserver.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	webserver.Run(":3000")
}
