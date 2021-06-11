package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectSql() *sql.DB {
	sqldb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/users")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	return sqldb
}
