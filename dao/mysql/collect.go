package mysql

import (
	"cyber-events-tracker/model"
	"database/sql"
)

func InsertCollectPaidMwSetEvent(param *model.ParamCollectPaidMwSetEvent) error {
	stmt, err := db.Prepare("INSERT INTO collect_paid_mw_set_events (chain_name, chain_id, block_number, tx_hash, namespace, profile_id, essence_id, total_supply, amount, recipient, currency, subscribe_required) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		param.ChainName,
		param.ChainID,
		param.BlockNumber,
		param.TxHash,
		param.Namespace,
		param.ProfileID,
		param.EssenceID,
		param.TotalSupply,
		param.Amount,
		param.Recipient,
		param.Currency,
		param.SubscribeRequired,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetCollectPaidMwSetEventParams(chainID uint64, profileID string, essenceID string) (*model.ParamCollectPaidMwSetEvent, error) {
	collectInfo := new(model.ParamCollectPaidMwSetEvent)
	sqlStr := "SELECT chain_name, chain_id, block_number, tx_hash, namespace, profile_id, essence_id, total_supply, amount, recipient, currency, subscribe_required FROM collect_paid_mw_set_events WHERE chain_id = ? AND profile_id = ? AND essence_id = ?"
	err := db.Get(collectInfo, sqlStr, chainID, profileID, essenceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return collectInfo, nil
}

func GetLatestTrackedCollectPaidMwSetBlockNumber(chainID uint64) (uint64, error) {
	collectInfo := new(model.ParamCollectPaidMwSetEvent)
	sqlStr := "SELECT block_number FROM collect_paid_mw_set_events WHERE chain_id = ? ORDER BY id DESC LIMIT 1"
	err := db.Get(collectInfo, sqlStr, chainID)
	if err != nil {
		return 0, err
	}
	return collectInfo.BlockNumber, nil
}
