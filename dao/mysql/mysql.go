package mysql

import (
	"cyber-events-tracker/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(mysqlConfig *settings.MySqlConfig) (err error) {
	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB)
	if db, err = sqlx.Connect("mysql", connectUrl); err != nil {
		return
	}

	db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
