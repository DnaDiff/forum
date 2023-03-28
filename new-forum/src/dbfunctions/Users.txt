package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	JoinedAt     time.Time `json:"joined"`
	PostCount    int       `json:"post_count"`
	CommentCount int       `json:"comment_count"`
}

// --------------------------------------------User functions--------------------------------------------

func CreateUser(db *sql.DB, username string, passwordHash string, email string) error {
	_, err := db.Exec("INSERT INTO users (username, passwrd, email) VALUES (?, ?, ?)", username, passwordHash, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(db *sql.DB, userID int) (*User, error) {
	var user User

	err := db.QueryRow("SELECT id, username, email, passwrd, joined, post_count, comment_count FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.JoinedAt, &user.PostCount, &user.CommentCount)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByName(db *sql.DB, username string) (*User, error) {
	var user User

	err := db.QueryRow("SELECT id, username, email, passwrd, joined, post_count, comment_count FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.JoinedAt, &user.PostCount, &user.CommentCount)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
