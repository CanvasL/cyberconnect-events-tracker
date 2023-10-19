package listener

import (
	"context"
	"log"
	"math/big"
	"cyber-events-tracker/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var topicCollectPaidMwSet = crypto.Keccak256Hash([]byte("CollectPaidMwSet(address,uint256,uint256,uint256,uint256,address,address,bool)"))

func CollectPaidMwEventListener(rpcUrl string, contractAddress common.Address) {
	ethClient, err := utils.GetEthClient(rpcUrl)
	if(err != nil) {
		return
	}
	currentBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCollectPaidMwSet}},
		FromBlock: big.NewInt(int64(currentBlockNumber)),
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("topicCollectPaidMwSet error happened:", err)
		case vLog := <-logs:
			log.Println("topicCollectPaidMwSet event received", vLog) // pointer to event log
			
		}
	}
}
