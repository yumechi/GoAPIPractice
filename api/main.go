package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/yumechi/GoAPIPractice/api/controller"
)

func main() {
	router := gin.Default()

	router.GET(
		"/pri",
		controller.PriGetting,
	)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		_ = router.Run(":37301")
	} else {
		_ = router.Run(":" + port)
	}
}