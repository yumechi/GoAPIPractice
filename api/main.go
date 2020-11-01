package main

import (
	"github.com/gin-gonic/gin"
	"local.packages/api/controller"
	"os"
)

func main() {
	router := gin.Default()

	router.GET(
		"/pri",
		controller.PriGetting,
	)
	router.POST(
		"/pri",
		controller.PriPosting,
	)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		_ = router.Run(":37301")
	} else {
		_ = router.Run(":" + port)
	}
}