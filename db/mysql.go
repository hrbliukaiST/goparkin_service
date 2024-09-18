go
package db

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "myproject/config"
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