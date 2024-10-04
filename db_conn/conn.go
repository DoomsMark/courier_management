package db_conn

import (
	"os"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"encoding/json"
)

type cfg struct{
	User string
	Passwd string
	Net string
	Addr string
	DBName string
}

func Connection() (*sql.DB, error){
	var db *sql.DB 
	var err error
	content, err := os.ReadFile("./config.json")
    if err != nil {
        return nil, err
    }
    var cfg_file cfg
    err = json.Unmarshal(content, &cfg_file)
    if err != nil {
        return nil, err
    }
	config := mysql.Config{
		User: cfg_file.User,
		Passwd: cfg_file.Passwd,
		Net: cfg_file.Net,
		Addr: cfg_file.Addr,
		DBName: cfg_file.DBName}

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil{
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil{
		return nil, err
	}
	return db, nil

}