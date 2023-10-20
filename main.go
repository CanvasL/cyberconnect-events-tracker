package main

import (
	"cyber-events-tracker/contract_abi"
	"cyber-events-tracker/dao/mysql"
	"cyber-events-tracker/listener"
	"cyber-events-tracker/router"
	"cyber-events-tracker/settings"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	if err := settings.Init(); err != nil {
		log.Fatalln("Init settings failed, ", err)
		return
	}

	if err := contract_abi.Init(); err != nil {
		log.Fatalln("Init contract abi failed, ", err)
		return
	}
	
	if err := mysql.Init(settings.Config.MySql); err != nil {
		fmt.Println("Init mysql failed, ", err)
		return
	}
	defer mysql.Close()

	go listener.CollectPaidMwSetEventListener(
		97, 
		common.HexToAddress(settings.Config.Contracts.BSCT.CollectPaidMw),
	)

	go listener.CreateProfileEventListener(
		97, 
		common.HexToAddress(settings.Config.Contracts.BSCT.ProfileNFT),
	)

	go listener.CollectPaidMwSetEventListener(
		56, 
		common.HexToAddress(settings.Config.Contracts.BSC.CollectPaidMw),
	)

	go listener.CreateProfileEventListener(
		56, 
		common.HexToAddress(settings.Config.Contracts.BSC.ProfileNFT),
	)

	r := router.SetupRouter(settings.Config.Mode)
	if err := r.Run(fmt.Sprintf(":%d", settings.Config.Port)); err != nil {
		log.Fatalln("Run server failed, ", err)
		return
	}
}
