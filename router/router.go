package router

import (
	"cyber-events-tracker/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1.GET("/collect", controller.GetCollectInfoHandler)

	v1.GET("/profile", controller.GetProfilesHandler)

	return r
}
