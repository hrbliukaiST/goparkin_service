package db

import (
	"database/sql"
	"goparkin_service/config"
	"log"

	// Add this line to import the time package

	"time"

	_ "github.com/go-sql-driver/mysql"
	// Add this import for http package
)

var MySQLDB *sql.DB

func InitMySQL() error {
	var err error
	for i := 0; i < 3; i++ {
		MySQLDB, err = sql.Open("mysql",config.AppConfig.MySQLDSN )
		if err == nil {
			// Attempt to ping the database to check connection
			if err = MySQLDB.Ping(); err == nil {
				return nil // Successfully connected
			}
		}
		log.Println("DB Server error: ", err.Error())
		time.Sleep(2 * time.Second) // Sleep before retrying
	}
	return err // Return the last error if all attempts fail
}