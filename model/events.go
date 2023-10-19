package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CreateProfileEvent struct {
	To        common.Address `json:"to"`
	ProfileID *big.Int       `json:"profileId"`
	Handle    string         `json:"handle"`
	Avatar    string         `json:"avatar"`
	Metadata  string         `json:"metadata"`
}

type CollectPaidMwEvent struct {
	Namespace         common.Address `json:"namespace"`
	ProfileID         *big.Int       `json:"profileId"`
	EssenceID         *big.Int       `json:"essenceId"`
	TotalSupply       *big.Int       `json:"totalSupply"`
	Amount            *big.Int       `json:"amount"`
	Recipient         common.Address `json:"recipient"`
	Currency          common.Address `json:"currency"`
	SubscribeRequired bool           `json:"subscribeRequired"`
}
