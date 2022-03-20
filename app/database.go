package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDB() *sql.DB {
	conf := GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8&parseTime=True"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
