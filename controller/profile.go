package controller

import "github.com/gin-gonic/gin"

func GetProfilesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ProfileCreated event emit",
	})
}
