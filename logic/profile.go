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

func SetProfilesInfo(chainID uint64, vLog types.Log) (err error) {
	eventData := struct {
		Handle   string
		Avatar   string
		Metadata string
	}{}
	err = contract_abi.AbiProfileNFT.UnpackIntoInterface(&eventData, "CreateProfile", vLog.Data)
	if err != nil {
		return
	}

	chainName, err := utils.GetChainName(chainID)
	if err != nil {
		return
	}

	param := &model.ParamCreateProfileEvent{
		ChainName:   chainName,
		ChainID:     chainID,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Hex(),
		To:          common.BytesToAddress(vLog.Topics[1].Bytes()).String(),
		ProfileID:   new(big.Int).SetBytes(vLog.Topics[2].Bytes()).String(),
		Handle:      eventData.Handle,
		Avatar:      eventData.Avatar,
		Metadata:    eventData.Metadata,
	}

	return mysql.InsertCreateProfileEvent(param)
}

func GetProfilesInfo(chainID uint64, account string) ([]*model.ParamCreateProfileEvent, error) {
	return mysql.GetCreateProfileEventParams(chainID, account)
}

func GetPreviousTrackedCreateProfileBlockNumber(chainID uint64) (uint64, error) {
	return mysql.GetLatestTrackedCreateProfileBlockNumber(chainID)
}