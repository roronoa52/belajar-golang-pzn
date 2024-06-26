package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.IfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}