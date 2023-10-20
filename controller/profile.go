package controller

import (
	"cyber-events-tracker/logic"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProfilesInfoHandler(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chain"), 10, 64)
	if(err != nil) {
		log.Fatalln("Parse chainID failed,", err)
		c.JSON(200, gin.H{
			"error": err,
			"message": nil,
		})
	}
	profileInfoList, err := logic.GetProfilesInfo(chainID, c.Param("account"))
	if err != nil {
		log.Fatalln("GetProfiles failed,", err)
		c.JSON(200, gin.H{
			"error": err,
			"message": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": profileInfoList,
			"error": nil,
		})
	}
}
