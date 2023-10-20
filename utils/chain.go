package utils

import (
	"errors"
	"os"
)

type ChainInfo struct {
	chainID   uint64
	ChainName string
}

var chainInfos = []ChainInfo{
	{chainID: 97, ChainName: "BSCT"},
	{chainID: 56, ChainName: "BSC"},
}

func GetChainName(chainID uint64) (string, error) {
	for _, chainInfo := range chainInfos {
		if chainInfo.chainID == chainID {
			return chainInfo.ChainName, nil
		}
	}
	return "", errors.New("ChainName not found for the given chainID")
}

func GetChainRPC(chainID uint64) (string) {
	if(chainID == 97) {
		return os.Getenv("BSCT_RPC_URL")
	}
	if(chainID == 56) {
		return os.Getenv("BSC_RPC_URL")
	}
	return ""
}