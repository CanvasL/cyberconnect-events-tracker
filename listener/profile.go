package listener

import (
	"context"
	"cyber-events-tracker/logic"
	"cyber-events-tracker/settings"
	"cyber-events-tracker/utils"
	"database/sql"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var topicCreateProfile = crypto.Keccak256Hash([]byte("CreateProfile(address,uint256,string,string,string)"))

func CreateProfileEventListener(chainID uint64, contractAddress common.Address) {
	ethClient, err := utils.GetEthClient(utils.GetChainRPC(chainID))
	if err != nil {
		log.Fatalf("[%d]: GetEthClient failed, %v", chainID, err)
		return
	}

	_currentBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("[%d]: Get current BlockNumber failed, %v", chainID, err)
		return
	}

	currentBlockNumber := big.NewInt(int64(_currentBlockNumber))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicCreateProfile}},
		FromBlock: currentBlockNumber,
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("[%s]: SubscribeFilterLogs CreateProfile failed, %v", chainID, err)
	}
	log.Printf("[%d]: Chan CreateProfile started.", chainID)

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("[%d]: Chan CreateProfile received error: %v", chainID, err)

		case vLog := <-logs:
			log.Printf("[%d]: Chan CreateProfile received vLog.", chainID)
			err = logic.SetProfilesInfo(chainID, vLog)
			if err != nil {
				log.Fatalf("[%d]: SetProfileInfo failed, %v", chainID, err)
			} else {
				log.Printf("[%d]: SetProfileInfo successfully.", chainID)
			}
		}
	}
}

func QueryCreateProfileEvents(chainID uint64, contractSetting *settings.ContractConfig) {
	if contractSetting.QueryHistory {
		ethClient, err := utils.GetEthClient(utils.GetChainRPC(chainID))
		if err != nil {
			log.Fatalf("[%d]: GetEthClient failed, %v", chainID, err)
			return
		}

		var _startAt *big.Int
		var _endAt *big.Int = big.NewInt(0)
		previousAt, err := logic.GetPreviousTrackedCreateProfileBlockNumber(chainID)
		if err != nil {
			if err == sql.ErrNoRows {
				_startAt = big.NewInt(contractSetting.StartAt)
				log.Printf("[%d]: Start query CreateProfile events at [%d]...", chainID, _startAt.Uint64())
			} else {
				log.Fatalf("[%d]: GetPreviousTrackedCreateProfileBlockNumber failed, %v", chainID, err)
				return
			}
		} else {
			_startAt = big.NewInt(int64(previousAt))
			log.Printf("[%d]: Continue query CreateProfile events at [%d]...", chainID, _startAt.Uint64())
		}

		var _currentBlockNumber uint64
		var currentBlockNumber *big.Int

		_currentBlockNumber, err = ethClient.BlockNumber(context.Background())
		if err != nil {
			log.Fatalf("[%d]: Get current BlockNumber failed, %v", chainID, err)
			return
		}
		currentBlockNumber = big.NewInt(int64(_currentBlockNumber))

		for {
			_endAt.Add(_startAt, big.NewInt(contractSetting.QueryInterval))
			if _endAt.Cmp(currentBlockNumber) > 0 {
				_endAt.Set(currentBlockNumber)
			}

			query := ethereum.FilterQuery{
				Addresses: []common.Address{common.HexToAddress(contractSetting.Address)},
				Topics:    [][]common.Hash{{topicCreateProfile}},
				FromBlock: _startAt,
				ToBlock:   _endAt,
			}

			historyLogs, err := ethClient.FilterLogs(context.Background(), query)
			if err != nil {
				log.Fatalf("[%d]: FilterLogs CreateProfile failed, %v", chainID, err)
				return
			}

			for _, historyLog := range historyLogs {
				err = logic.SetProfilesInfo(chainID, historyLog)
				if err != nil {
					log.Fatalf("[%d]: SetProfileInfo failed, %v", chainID, err)
				}
			}

			time.Sleep(200 * time.Millisecond)

			_currentBlockNumber, err = ethClient.BlockNumber(context.Background())
			if err != nil {
				log.Fatalf("[%d]: Get current BlockNumber failed, %v", chainID, err)
			} else {
				currentBlockNumber = big.NewInt(int64(_currentBlockNumber))
			}

			if _endAt.Cmp(currentBlockNumber) == 0 {
				break
			}

			_startAt.Add(_endAt, big.NewInt(1))
		}

		log.Printf("[%d]: Query CreateProfile history events successfully.", chainID)
	}
}
