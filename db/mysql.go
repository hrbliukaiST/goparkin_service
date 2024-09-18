package db

import (
	"database/sql"
	"log"

	"goparkin_service/config"

	_ "github.com/go-sql-driver/mysql"
)

var MySQLDB *sql.DB

func InitMySQL() {
    var err error
    MySQLDB, err = sql.Open("mysql", config.AppConfig.MySQLDSN)
    if err != nil {
        log.Fatal(err)
    }
    if err := MySQLDB.Ping(); err != nil {
        log.Fatal(err)
    }
}