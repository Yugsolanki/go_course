package main

import (
	"database/sql"
	"fmt"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func initializeDB() (*sql.DB, error) {
	once.Do(func() {
		fmt.Println("Initializing DB...")
		db, err = sql.Open("postgres", "connection-string-here")
	})
	return db, err
}

func main() {
	conn1, _ := initializeDB()
	conn2, _ := initializeDB()

	fmt.Println(conn1, conn2)
}
