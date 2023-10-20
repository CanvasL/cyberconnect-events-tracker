package model

type ParamCreateProfileEvent struct {
	ChainName   string `json:"chain_name" db:"chain_name"`
	ChainID     uint64 `json:"chain_id" db:"chain_id"`
	BlockNumber uint64 `json:"block_number" db:"block_number"`
	TxHash      string `json:"tx_hash" db:"tx_hash"`
	To          string `json:"to" db:"to"`
	ProfileID   string `json:"profile_id" db:"profile_id"`
	Handle      string `json:"handle" db:"handle"`
	Avatar      string `json:"avatar" db:"avatar"`
	Metadata    string `json:"metadata" db:"metadata"`
}

type ParamCollectPaidMwSetEvent struct {
	ChainName         string `json:"chain_name" db:"chain_name"`
	ChainID           uint64 `json:"chain_id" db:"chain_id"`
	BlockNumber       uint64 `json:"block_number" db:"block_number"`
	TxHash            string `json:"tx_hash" db:"tx_hash"`
	Namespace         string `json:"namespace" db:"namespace"`
	ProfileID         string `json:"profile_id" db:"profile_id"`
	EssenceID         string `json:"essence_id" db:"essence_id"`
	TotalSupply       string `json:"total_supply" db:"total_supply"`
	Amount            string `json:"amount" db:"amount"`
	Recipient         string `json:"recipient" db:"recipient"`
	Currency          string `json:"currency" db:"currency"`
	SubscribeRequired bool   `json:"subscribe_required" db:"subscribe_required"`
}
