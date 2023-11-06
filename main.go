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
		log.Println("Init mysql failed, ", err)
		return
	}
	defer mysql.Close()

	log.Printf("\033[1;34m╔══════════════════════════════════════════╗\033[0m")
	log.Printf("\033[1;34m║        Cyber Events Tracker %s       ║\033[0m", settings.Config.Version)
	log.Printf("\033[1;34m╚══════════════════════════════════════════╝\033[0m")

	listener.QueryCollectPaidMwSetEvents(
		97,
		settings.Config.Contracts.BSCT.CollectPaidMw,
	)

	listener.QueryCreateProfileEvents(
		97,
		settings.Config.Contracts.BSCT.ProfileNFT,
	)

	listener.QueryCollectPaidMwSetEvents(
		56,
		settings.Config.Contracts.BSC.CollectPaidMw,
	)

	listener.QueryCreateProfileEvents(
		56,
		settings.Config.Contracts.BSC.ProfileNFT,
	)

	go listener.CollectPaidMwSetEventListener(
		97,
		common.HexToAddress(settings.Config.Contracts.BSCT.CollectPaidMw.Address),
	)

	go listener.CreateProfileEventListener(
		97,
		common.HexToAddress(settings.Config.Contracts.BSCT.ProfileNFT.Address),
	)

	go listener.CollectPaidMwSetEventListener(
		56,
		common.HexToAddress(settings.Config.Contracts.BSC.CollectPaidMw.Address),
	)

	go listener.CreateProfileEventListener(
		56,
		common.HexToAddress(settings.Config.Contracts.BSC.ProfileNFT.Address),
	)

	r := router.SetupRouter(settings.Config.Mode)
	if err := r.Run(fmt.Sprintf(":%d", settings.Config.Port)); err != nil {
		log.Fatalln("Run server failed, ", err)
		return
	}
}
