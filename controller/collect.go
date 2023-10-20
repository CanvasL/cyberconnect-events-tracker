package controller

import (
	"cyber-events-tracker/logic"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCollectInfoHandler(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chain"), 10, 64)
	if(err != nil) {
		log.Fatalln("Parse chainID failed,", err)
		c.JSON(200, gin.H{
			"error": err,
			"message": nil,
		})
	}
	collectInfo, err := logic.GetCollectInfo(chainID, c.Param("profile"), c.Param("essence"))
	if err != nil {
		log.Fatalln("GetCollectInfo failed,", err)
		c.JSON(200, gin.H{
			"error": err,
			"message": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": collectInfo,
			"error": nil,
		})
	}
}