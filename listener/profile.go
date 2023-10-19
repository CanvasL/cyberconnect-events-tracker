package listener

import (
	"context"
	"cyber-events-tracker/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var topicCreateProfile = crypto.Keccak256Hash([]byte("CreateProfile(address,uint256,string,string,string)"))

func CreateProfileEventListener(rpcUrl string, contractAddress common.Address) {
	ethClient, err := utils.GetEthClient(rpcUrl)
	if(err != nil) {
		return
	}

	blockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCreateProfile}},
		FromBlock: big.NewInt(int64(blockNumber - 10000)),
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("topicCreateProfile error happened:", err)
		case vLog := <-logs:
			log.Println("topicCreateProfile event received", vLog) // pointer to event log
		}
	}
}
