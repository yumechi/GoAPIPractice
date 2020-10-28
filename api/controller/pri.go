package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PriGetting(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "puripuripuri~~~~",
	})
}