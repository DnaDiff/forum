package database

import (
	"database/sql"
	"log"
	"os"
)

// InitDatabase initializes the database
func InitDatabase(db *sql.DB) {
	sqlScript, err := os.ReadFile("./database/init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatal(err)
	}
}
