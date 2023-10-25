package contract_abi

import (
	_ "embed"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var AbiCollectPaidMw abi.ABI
var AbiProfileNFT abi.ABI

var (
	//go:embed abi/CollectPaidMw.json
	CollectPaidMw string
	//go:embed abi/ProfileNFT.json
	ProfileNFT string
)

func Init() (err error) {
	AbiCollectPaidMw, err = abi.JSON(strings.NewReader(CollectPaidMw))
	if err != nil {
		return
	}

	AbiProfileNFT, err = abi.JSON(strings.NewReader(ProfileNFT))
	if err != nil {
		return
	}
	return
}
