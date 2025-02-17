package main

import (
	"addressintelligence/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/v1/get_label", router.GetLabel)

	r.Run(":7777")
}

// ok
