package main

import (
	"cyber-events-tracker/abi_reader"
	"cyber-events-tracker/dao/mysql"
	"cyber-events-tracker/listener"
	"cyber-events-tracker/router"
	"cyber-events-tracker/settings"
	"fmt"
	"log"
	"os"

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

	go listener.CollectPaidMwEventListener(os.Getenv("BSCT_RPC_URL"), common.HexToAddress(settings.Config.Contracts.BSCT.CollectPaidMw))

	go listener.CreateProfileEventListener(os.Getenv("BSCT_RPC_URL"), common.HexToAddress(settings.Config.Contracts.BSCT.ProfileNFT))

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Config.Port)); err != nil {
		log.Fatalln("run server failed.", err)
		return
	}

	// // Define MySQL database connection
	// cfg := mysql.NewConfig()
	// cfg.User = "username"
	// cfg.Passwd = "password"
	// cfg.Net = "tcp"
	// cfg.Addr = "localhost:3306"
	// cfg.DBName = "database"
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // Connect to Ethereum network
	// client, err := ethclient.Dial("https://rinkeby.infura.io/v3/your-project-id")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Define contract address and ABI
	// contractAddress := common.HexToAddress("0x123456...")
	// contractABI := "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"CreateProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"namespace\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"subscribeRequired\",\"type\":\"bool\"}],\"name\":\"CollectPaidMwSet\",\"type\":\"event\"}]"

	// // Instantiate contract instance
	// contract, err := NewMyContract(contractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Subscribe to CreateProfile event
	// createProfileEventCh := make(chan *MyContractCreateProfile)
	// createProfileSub, err := contract.WatchCreateProfile(&bind.WatchOpts{}, createProfileEventCh)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer createProfileSub.Unsubscribe()

	// // Subscribe to CollectPaidMwSet event
	// collectPaidMwEventCh := make(chan *MyContractCollectPaidMwSet)
	// collectPaidMwSub, err := contract.WatchCollectPaidMwSet(&bind.WatchOpts{}, collectPaidMwEventCh)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer collectPaidMwSub.Unsubscribe()

	// // Define event handlers
	// createProfileHandler := func(event *MyContractCreateProfile) {
	// 	fmt.Println("CreateProfile event received:", event.To.Hex(), event.ProfileID, event.Handle, event.Avatar, event.Metadata)

	// 	// Store Profile event data in MySQL database
	// 	_, err = db.Exec(`
	// 		INSERT INTO profile_events (blockchain_id, to_address, profile_id, handle, avatar, metadata)
	// 		VALUES (?, ?, ?, ?, ?, ?)
	// 	`, "rinkeby", event.To.Hex(), event.ProfileID.String(), event.Handle, event.Avatar, event.Metadata)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// collectPaidMwHandler := func(event *MyContractCollectPaidMwSet) {
	// 	fmt.Println("CollectPaidMwSet event received:", event.Namespace.Hex(), event.ProfileID, event.EssenceID, event.TotalSupply, event.Amount, event.Recipient.Hex(), event.Currency.Hex(), event.SubscribeRequired)

	// 	// Store CollectPaidMw event data in MySQL database
	// 	_, err = db.Exec(`
	// 		INSERT INTO collect_paid_mw_events (blockchain_id, namespace, profile_id, essence_id, total_supply, amount, recipient, currency, subscribe_required)
	// 		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	// 	`, "rinkeby", event.Namespace.Hex(), event.ProfileID.String(), event.EssenceID.String(), event.TotalSupply.String(), event.Amount.String(), event.Recipient.Hex(), event.Currency.Hex(), event.SubscribeRequired)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // Start event loop
	// ctx := context.Background()
	// subCtx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// createProfileSubErrCh := make(chan error)
	// go func() {
	// 	createProfileSubErrCh <- event.ReadLoop(subCtx, createProfileSub)
	// }()

	// collectPaidMwSubErrCh := make(chan error)
	// go func() {
	// 	collectPaidMwSubErrCh <- event.ReadLoop(subCtx, collectPaidMwSub)
	// }()

	// // Wait for events to be received
	// select {
	// case err := <-createProfileSubErrCh:
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case err := <-collectPaidMwSubErrCh:
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
