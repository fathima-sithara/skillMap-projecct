package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("skill map project")

	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "skill-map-home-page ",
		})
	})
	r.Run()
}

