package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthClient(rpcUrl string) (*ethclient.Client, error) {
	if ethClient, err := ethclient.Dial(rpcUrl); err != nil {
		log.Fatalln("init ethclient failed, ", err)
		return nil, err
	} else {
		return ethClient, nil
	}
}
