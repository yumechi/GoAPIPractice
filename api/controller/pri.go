package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Idol struct {
	Name string `json:"name"`
	Age int `json:"age"`
	// TODO: デフォルト値が設定できてない
	Greet string `json:"greet,default=puripuri~"`
}

func PriGetting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "puripuripuri~~~~",
	})
}

func PriPosting(c *gin.Context) {
	var json Idol
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{
		"name": json.Name,
		"Age": json.Age,
		"greet": json.Greet,
	})
}