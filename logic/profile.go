package logic

import (
	// "cyber-events-tracker/abi_reader"
	"cyber-events-tracker/abi_reader"
	"cyber-events-tracker/dao/mysql"
	"cyber-events-tracker/model"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SetProfileInfo(vLog types.Log) (err error) {
	eventData := struct {
		Handle   string
		Avatar   string
		Metadata string
	}{}
	err = abi_reader.AbiProfileNFT.UnpackIntoInterface(&eventData, "CreateProfile", vLog.Data)
	if err != nil {
		return
	}
	// log.Println("SetProfileInfo")

	// param := &model.ParamCreateProfileEvent{
	// 	ChainName:   "BSCT",
	// 	ChainId:     97,
	// 	BlockNumber: 34339237,
	// 	TxHash: common.HexToHash("0x6237176545160fea5c49d4bd27ce797597c5771dd3302056b35a82677ee3924d"),
	// 	To:          common.BytesToAddress([]byte("0x13a6D1fe418de7e5B03Fb4a15352DfeA3249eAA4")),
	// 	ProfileID:   big.NewInt(int64(2240)),
	// 	Handle:      "profile1697714957192",
	// 	Avatar:      "https://testnet.bscscan.com/assets/bsc/images/svg/logos/logo-bscscan.svg?v=23.10.2.0",
	// 	Metadata:    "",
	// }
	param := &model.ParamCreateProfileEvent{
		ChainName:   "BSCT",
		ChainId:     97,
		BlockNumber: vLog.BlockNumber,
		TxHash: vLog.TxHash,
		To:          common.BytesToAddress(vLog.Topics[1].Bytes()),
		ProfileID:   new(big.Int).SetBytes(vLog.Topics[2].Bytes()),
		Handle:      eventData.Handle,
		Avatar:      eventData.Avatar,
		Metadata:    eventData.Metadata,
	}

	return mysql.InsertCreateProfileEvent(param)
}

func GetProfiles() (err error) {
	return nil
}
