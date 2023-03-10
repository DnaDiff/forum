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

	fmt.Println("User created successfully")

	return nil
}

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
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return u, nil
}
