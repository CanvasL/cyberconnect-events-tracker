package utils

import "github.com/ethereum/go-ethereum/ethclient"


func GetEthClient(rpcUrl string) (*ethclient.Client, error) {
	if ethClient, err := ethclient.Dial(rpcUrl); err != nil {
		return nil, err
	} else {
		return ethClient, nil
	}
}
