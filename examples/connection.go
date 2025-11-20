package examples

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/testings")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
