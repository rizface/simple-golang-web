package helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pbl-orkom/app"
	"time"
)

func Connection() *sql.DB {
	db, err := sql.Open(app.DATABASE_DRIVER, "root:root@tcp(localhost:3306)/pbl_orkom?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db
}
