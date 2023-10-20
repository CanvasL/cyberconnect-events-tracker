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

var topicCreateProfile = crypto.Keccak256Hash([]byte("CreateProfile(address,uint256,string,string,string)"))

func CreateProfileEventListener(chainID uint64, contractAddress common.Address) {
	ethClient, err := utils.GetEthClient(utils.GetChainRPC(chainID))
	if(err != nil) {
		log.Fatalf("[%d]: GetEthClient failed, %v", chainID, err)
		return
	}

	currentBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("[%d]: Get current BlockNumber failed, %v", chainID, err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCreateProfile}},
		FromBlock: big.NewInt(int64(currentBlockNumber)),
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("[%s]: SubscribeFilterLogs failed, %v", chainID, err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("[%d]: Chan CollectPaidMwSet received error: %v", chainID, err)

		case vLog := <-logs:
			log.Printf("[%d]: Chan CollectPaidMwSet received vLog.", chainID) 
			err = logic.SetProfilesInfo(chainID, vLog)
			if(err != nil) {
				log.Fatalf("[%d]: SetCollectInfo failed, %v", chainID, err)
			} else {
				log.Printf("[%d]: SetCollectInfo successfully.", chainID)
			}
		}
	}
}
