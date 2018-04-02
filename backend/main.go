package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(
		gin.Recovery(),
	)

	api := router.Group("/api")
	{
		api.GET("/me", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world!")
		})
	}

	router.Run(":6066")
}
