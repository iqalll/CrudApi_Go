package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil")
	}

	return db, nil
}
