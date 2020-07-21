package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8088"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
