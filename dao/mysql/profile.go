package mysql

import (
	"cyber-events-tracker/model"
	"fmt"
	"log"
)

func InsertCreateProfileEvent(param *model.ParamCreateProfileEvent) (err error) {
	log.Println("mysql param:", param)
	// 准备插入语句
	stmt, err := db.Prepare("INSERT INTO create_profile_events (chain_name, chain_id, block_number, tx_hash, `to`, profile_id, handle, avatar, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(param.ChainName, param.ChainId, param.BlockNumber, param.TxHash.Hex(), param.To.Hex(), param.ProfileID.String(), param.Handle, param.Avatar, param.Metadata)
	if err != nil {
		return err
	}

	fmt.Println("CreateProfileEvent inserted successfully")
	return nil
}
