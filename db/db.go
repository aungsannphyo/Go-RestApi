package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBInstance *sql.DB
	once       sync.Once
)

func InitDB() *sql.DB {
	once.Do(func() {
		var err error
		DBInstance, err = sql.Open("sqlite3", "api.db")
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		DBInstance.SetMaxOpenConns(10)
		DBInstance.SetMaxIdleConns(5)
		createTable()
	})
	return DBInstance
}
