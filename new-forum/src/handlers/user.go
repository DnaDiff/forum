package handlers

import (
	"database/sql"
	"net/http"
)

type User struct {
	ID       string `json:"ID"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const DEFAULT_AVATAR = "https://st3.depositphotos.com/6672868/13701/v/600/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"

// Placeholder data
var users = map[string]User{
	"123456789": {ID: "123456789", Avatar: DEFAULT_AVATAR, Username: "John_Doe", Password: "password", Email: "john@doe.com"},
	"234567890": {ID: "234567890", Avatar: DEFAULT_AVATAR, Username: "janedoe", Password: "password", Email: "jane@doe.com"},
	"345678901": {ID: "345678901", Avatar: DEFAULT_AVATAR, Username: "partyboi", Password: "password", Email: "partyboi@forever.com"},
}

func getUser(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) User {
	// Fetch user from database below

	// Placeholder data
	if user, ok := users[userID]; ok {
		return user
	}
	return User{}
}
