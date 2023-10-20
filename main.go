package main

import (
	"cyber-events-tracker/abi_reader"
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
		log.Fatalln("init settings failed.", err)
		return
	}

	if err := abi_reader.Init(); err != nil {
		log.Fatalln("init abi failed.", err)
		return
	}
	
	if err := mysql.Init(settings.Config.MySql); err != nil {
		fmt.Println("init mysql failed, err:", err)
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

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Config.Port)); err != nil {
		log.Fatalln("run server failed.", err)
		return
	}
}
