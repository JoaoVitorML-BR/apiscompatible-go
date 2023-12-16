package database

import (
	"database/sql"
	"login/infra/tools/database/config"

	_ "github.com/go-sql-driver/mysql"
)

// OPen connection with database
func DatabaseAPI() (*sql.DB, error){
	config.LoadInfos()

	urlConnect := config.DatabaseConnectString
	db, err := sql.Open("mysql", urlConnect)
	if err != nil{
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}

	return db, nil
}