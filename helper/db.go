package helper

import (
	"database/sql"
	"os"
	"pbl-orkom/app"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	dsn := "root:root@tcp("+os.Getenv("MYSQL_HOST")+":3306)/pbl_orkom?parseTime=true"
	db, err := sql.Open(app.DATABASE_DRIVER,dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db
}
