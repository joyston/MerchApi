package dbfunctions

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBConnect() (bool, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return false, fmt.Errorf("DBConnect: error at sql.Open %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return false, fmt.Errorf("DBConnect: error at db.Ping %v", pingErr)
	}
	return true, nil
}
