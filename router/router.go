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

	v1.GET("/:chain/collect-info/:profile/:essence", controller.GetCollectInfoHandler)

	v1.GET("/:chain/profiles-info/:account", controller.GetProfilesInfoHandler)

	return r
}
