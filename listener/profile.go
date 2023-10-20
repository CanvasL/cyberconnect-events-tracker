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
		log.Fatalln("GetEthClient failed, ", err)
		return
	}

	currentBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Get current BlockNumber failed, ", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCreateProfile}},
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
			log.Fatal("Chan CreateProfile received error:", err)
		case vLog := <-logs:
			log.Println("Chan CreateProfile received vLog.") 
			err = logic.SetProfilesInfo(chainID, vLog)
			if(err != nil) {
				log.Fatalln("SetProfilesInfo failed, ", err)
			} else {
				log.Println("SetProfilesInfo successfully.")
			}
		}
	}
}
