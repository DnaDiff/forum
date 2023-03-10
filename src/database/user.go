package database

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID             int
	ProfilePicture string
	Username       string
	Age            int
	Gender         string
	FirstName      string
	LastName       string
	Password       string
	Email          string
	Joined         string
	PostCount      int
	CommentCount   int
	LikeCount      int
	DislikeCount   int
}

// CreateUser creates a new user in the database, if the profile picture is empty, a default image will be used
func CreateUser(db *sql.DB, u *User) error {
	query := "INSERT INTO users(username, age, gender, first_name, last_name, passwrd, email"
	if u.ProfilePicture != "" {
		query += ", profile_picture"
	}
	query += ") VALUES (?, ?, ?, ?, ?, ?, ?"
	if u.ProfilePicture != "" {
		query += ", ?"
	}
	query += ")"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	args := []interface{}{u.Username, u.Age, u.Gender, u.FirstName, u.LastName, u.Password, u.Email}
	if u.ProfilePicture != "" {
		args = append(args, u.ProfilePicture)
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("execute statement error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows error: %w", err)
	}

	if rowsAffected != 1 {
		return fmt.Errorf("unexpected rows affected: %d", rowsAffected)
	}

	fmt.Println("User: " + "\"" + u.Username + "\"" + " created successfully")

	return nil
}

// GetUserByUsername returns a user from the database by username
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT id, profile_picture, username, age, gender, first_name, last_name, passwrd, email, joined, post_count, comment_count, like_count, dislike_count FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	u := &User{}
	err = row.Scan(&u.ID, &u.ProfilePicture, &u.Username, &u.Age, &u.Gender, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Joined, &u.PostCount, &u.CommentCount, &u.LikeCount, &u.DislikeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user: " + "\"" + username + "\"" + " not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return u, nil
}

// DeleteUserByUsername deletes a user from the database by username
func DeleteUserByUsername(db *sql.DB, username string) error {
	query := "DELETE FROM users WHERE username = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(username)
	if err != nil {
		return fmt.Errorf("execute statement error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows error: %w", err)
	}

	if rowsAffected != 1 {
		return fmt.Errorf("unexpected rows affected: %d", rowsAffected)
	}

	fmt.Println("User: " + "\"" + username + "\"" + " deleted successfully")

	return nil
}

// CheckDuplicateUser checks if the username or email already exists in the database
func CheckDuplicateUser(db *sql.DB, username string, email string) bool {
	query := "SELECT count(*) FROM users WHERE username = ? OR email = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(username, email)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

// CheckDuplicateUsername checks if the username already exists in the database
func CheckDuplicateUsername(db *sql.DB, username string) bool {
	query := "SELECT count(*) FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

// CheckDuplicateEmail checks if the email is already in use
func CheckDuplicateEmail(db *sql.DB, email string) bool {
	query := "SELECT count(*) FROM users WHERE email = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
