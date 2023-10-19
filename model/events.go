package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ParamCreateProfileEvent struct {
	ChainName   string         `json:"chain_name"`
	ChainId     uint64         `json:"chain_id"`
	BlockNumber uint64         `json:"block_number"`
	TxHash      common.Hash    `json:"tx_hash"`
	To          common.Address `json:"to"`
	ProfileID   *big.Int       `json:"profile_id"`
	Handle      string         `json:"handle"`
	Avatar      string         `json:"avatar"`
	Metadata    string         `json:"metadata"`
}

type ParamCollectPaidMwSetEvent struct {
	ChainName         string         `json:"chain_name"`
	ChainId           string         `json:"chain_id"`
	BlockNumber       uint64         `json:"block_number"`
	TxHash            common.Hash    `json:"tx_hash"`
	Namespace         common.Address `json:"namespace"`
	ProfileID         *big.Int       `json:"profile_id"`
	EssenceID         *big.Int       `json:"essenceId"`
	TotalSupply       *big.Int       `json:"total_supply"`
	Amount            *big.Int       `json:"amount"`
	Recipient         common.Address `json:"recipient"`
	Currency          common.Address `json:"currency"`
	SubscribeRequired bool           `json:"subscribe_required"`
}
