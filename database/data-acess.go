package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func DBInit() {

	config := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Addr:                 os.Getenv("DATABASE_URL"),
		Net:                  "tcp",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	d, _ := sql.Open("mysql", config.FormatDSN())

	//最大開幾個連線
	d.SetMaxOpenConns(5)
	//最多幾個閒置連線
	d.SetMaxIdleConns(2)
	//閒置多久後刪除連線
	d.SetConnMaxLifetime(time.Hour)
	//Connection Test
	err := d.Ping()

	if err != nil {
		panic(err)
	}

	db = d
}
