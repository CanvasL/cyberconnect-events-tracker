package controller

import "github.com/gin-gonic/gin"

func GetCollectInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Collect Event Emit",
	})
}