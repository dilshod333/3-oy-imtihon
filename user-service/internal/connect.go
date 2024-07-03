package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" 
)


const (
	host   = "localhost"
	dbname = "anyy"
	dbport = 5432
	dbuser = "postgres"
	dbpass = "dilshod"
)

func Connect() (*sql.DB) {
	dbInfo := fmt.Sprintf("host=%s dbname=%s port=%d user=%s password=%s sslmode=disable",
		host, dbname, dbport, dbuser, dbpass)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Printf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("verifying error: %v", err)
	}

	return db
}
