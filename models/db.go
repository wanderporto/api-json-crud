package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER  = "mysql"
	DBNAME  = "api-json-go"
	USER_DB = "root"
	PORT    = "3306"
	HOST    = "127.0.0.1"
)

var password = os.Getenv("DB_PASSWORD")

func Connect() *sql.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER_DB, password, HOST, PORT, DBNAME)

	con, err := sql.Open(DRIVER, URL)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return con
}
