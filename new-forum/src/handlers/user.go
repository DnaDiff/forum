package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

type User struct {
	ID       string `json:"ID"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) User {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Printf("Error converting userID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return User{}
	}
	// Fetch user from database below
	userDB, err := database.GetUserByID(db, userIDInt)
	if err != nil {
		fmt.Printf("Error fetching user from database: %v\n", err)
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return User{}
	}
	user := User{
		ID:       strconv.Itoa(userDB.ID),
		Avatar:   userDB.ProfilePicture,
		Username: userDB.Username,
		Password: userDB.Password,
		Email:    userDB.Email,
	}
	return user
}
