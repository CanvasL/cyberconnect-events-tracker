package logic

import (
	"cyber-events-tracker/contract_abi"
	"cyber-events-tracker/dao/mysql"
	"cyber-events-tracker/model"
	"cyber-events-tracker/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SetCollectInfo(chainID uint64, vLog types.Log) (err error) {
	eventData := struct {
		TotalSupply       *big.Int
		Amount            *big.Int
		Recipient         common.Address
		Currency          common.Address
		SubscribeRequired bool
	}{}
	err = contract_abi.AbiCollectPaidMw.UnpackIntoInterface(&eventData, "CollectPaidMwSet", vLog.Data)
	if err != nil {
		return
	}

	chainName, err := utils.GetChainName(chainID)
	if err != nil {
		return
	}

	param := &model.ParamCollectPaidMwSetEvent{
		ChainName:         chainName,
		ChainID:           chainID,
		BlockNumber:       vLog.BlockNumber,
		TxHash:            vLog.TxHash.Hex(),
		Namespace:         common.BytesToAddress(vLog.Topics[1].Bytes()).String(),
		ProfileID:         new(big.Int).SetBytes(vLog.Topics[2].Bytes()).String(),
		EssenceID:         new(big.Int).SetBytes(vLog.Topics[3].Bytes()).String(),
		TotalSupply:       eventData.TotalSupply.String(),
		Amount:            eventData.Amount.String(),
		Recipient:         eventData.Recipient.String(),
		Currency:          eventData.Currency.String(),
		SubscribeRequired: eventData.SubscribeRequired,
	}

	return mysql.InsertCollectPaidMwSetEvent(param)
}

func GetCollectInfo(chainID uint64, profileID string, essenceID string) (*model.ParamCollectPaidMwSetEvent, error) {
	return mysql.GetCollectPaidMwSetEventParams(chainID, profileID, essenceID)
}
