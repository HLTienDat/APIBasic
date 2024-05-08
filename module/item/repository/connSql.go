package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "091372"
	dbname   = "Users"
)

var (
	database *sql.DB
	once     sync.Once
)

func ConnectDB() (*sql.DB, error) {
	once.Do(func() {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
		if err != nil {
			log.Fatal("Error connecting to the database: ", err)
		}
		database = db
	})
	return database, nil
}
