package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InitDB() *sqlx.DB {
	database, err := sqlx.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/policymanager")
	if err != nil {
		fmt.Println("open mysql failed,", err)

	}
	return database
}
