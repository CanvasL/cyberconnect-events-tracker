package listener

import (
	"context"
	"cyber-events-tracker/logic"
	"cyber-events-tracker/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var topicCollectPaidMwSet = crypto.Keccak256Hash([]byte("CollectPaidMwSet(address,uint256,uint256,uint256,uint256,address,address,bool)"))

func CollectPaidMwSetEventListener(chainID uint64, contractAddress common.Address) {
	ethClient, err := utils.GetEthClient(utils.GetChainRPC(chainID))
	if(err != nil) {
		log.Fatalln("GetEthClient failed, ", err)
		return
	}

	currentBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Get current BlockNumber failed, ", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCollectPaidMwSet}},
		FromBlock: big.NewInt(int64(currentBlockNumber)),
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("SubscribeFilterLogs failed, ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("Chan CollectPaidMwSet received error:", err)
		case vLog := <-logs:
			log.Println("Chan CollectPaidMwSet received vLog.") 
			err = logic.SetCollectInfo(chainID, vLog)
			if(err != nil) {
				log.Fatalln("SetCollectInfo failed, ", err)
			} else {
				log.Println("SetCollectInfo successfully.")
			}
		}
	}
}
