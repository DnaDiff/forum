package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func AddUser(db *sql.DB) {
	username := "John Doe"
	emails := "john.doe@example.com"
	password := "password"

	result, err := db.Exec("INSERT INTO users (username, email, passwrd) VALUES (?, ?, ?)", username, emails, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d user added.\n", result)

	var id int
	var name, email string

	err = db.QueryRow("SELECT id, username, email FROM users WHERE username = ?", username).Scan(&id, &name, &email)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{id, name, email}

	userInfo, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(userInfo) + "\n" + "User added successfully.")
}

func GetUser(db *sql.DB) (user struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Joined string `json:"joined"`
}, err error) {
	var id int
	var name, email, joined string

	err = db.QueryRow("SELECT id, username, email, joined FROM users WHERE id = ?", 1).Scan(&id, &name, &email, &joined)
	if err != nil {
		return user, err
	}

	user = struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Joined string `json:"joined"`
	}{id, name, email, joined}

	return user, nil
}
