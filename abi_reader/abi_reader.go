package abi_reader

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var AbiCollectPaidMw abi.ABI
var AbiProfileNFT abi.ABI

func Init() (err error) {
	path, _ := filepath.Abs("./abi_reader/CollectPaidMw.json")
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	AbiCollectPaidMw, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return
	}

	path, _ = filepath.Abs("./abi_reader/ProfileNFT.json")
	file, err = os.ReadFile(path)
	if err != nil {
		return
	}
	AbiProfileNFT, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return
	}
	return
}
