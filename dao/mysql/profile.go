package mysql

import (
	"cyber-events-tracker/model"
)

func InsertCreateProfileEvent(param *model.ParamCreateProfileEvent) (err error) {
	stmt, err := db.Prepare("INSERT INTO create_profile_events (chain_name, chain_id, block_number, tx_hash, `to`, profile_id, handle, avatar, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(param.ChainName, param.ChainID, param.BlockNumber, param.TxHash, param.To, param.ProfileID, param.Handle, param.Avatar, param.Metadata)
	if err != nil {
		return err
	}

	return nil
}

func GetCreateProfileEventParams(chainID uint64, account string) ([]*model.ParamCreateProfileEvent, error) {
	profileInfoList := make([]*model.ParamCreateProfileEvent, 0, 2)
	sqlStr := "SELECT chain_name, chain_id, block_number, tx_hash, `to`, profile_id, handle, avatar, metadata FROM create_profile_events WHERE chain_id = ? AND `to` = ? ORDER BY block_number DESC"
	err := db.Select(&profileInfoList, sqlStr, chainID, account)
	return profileInfoList, err
}

func GetLatestTrackedCreateProfileBlockNumber(chainID uint64) (uint64, error) {
	profileInfo := new(model.ParamCreateProfileEvent)
	sqlStr := "SELECT block_number FROM create_profile_events WHERE chain_id = ? ORDER BY id DESC LIMIT 1"
	err := db.Get(profileInfo, sqlStr, chainID)
	if err != nil {
		return 0, err
	}
	return profileInfo.BlockNumber, nil
}
