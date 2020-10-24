package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "puripuripuri~~~~",
		})
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		_ = r.Run(":37301")
	} else {
		_ = r.Run(":" + port)
	}
}
