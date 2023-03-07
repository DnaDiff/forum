package database

import (
	"database/sql"
	"io/ioutil"
	"log"
)

func InitDatabase(db *sql.DB) {
	// Execute the contents of the init.sql file
	sqlScript, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatal(err)
	}
}
