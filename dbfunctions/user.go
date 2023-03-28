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
}

// Working functions

// CreateUser creates a new user in the database, if the profile picture is empty, a default image will be used
func CreateUser(db *sql.DB, u *User) error {

	if len([]rune(u.Username)) > 12 {
		return fmt.Errorf("username is too long")
	}

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

// GetUserLikes gets the total number of likes the user has received on their posts and comments
func GetTotalUserLikes(db *sql.DB, userID int) (int, error) {
	query := `SELECT COUNT(*) 
              FROM likes 
              WHERE post_id IN (SELECT id FROM posts WHERE user_id = ?) 
                 OR comment_id IN (SELECT id FROM comments WHERE user_id = ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID, userID)

	var totalLikes int
	err = row.Scan(&totalLikes)
	if err != nil {
		return 0, fmt.Errorf("query error: %w", err)
	}

	return totalLikes, nil
}

// GetTotalUserDislikes gets the total number of dislikes the user has received on their posts and comments
func GetTotalUserDislikes(db *sql.DB, userID int) (int, error) {
	query := `SELECT COUNT(*) 
              FROM dislikes 
              WHERE post_id IN (SELECT id FROM posts WHERE user_id = ?) 
                 OR comment_id IN (SELECT id FROM comments WHERE user_id = ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID, userID)

	var totalDislikes int
	err = row.Scan(&totalDislikes)
	if err != nil {
		return 0, fmt.Errorf("query error: %w", err)
	}

	return totalDislikes, nil
}

// GetTotalUserPosts gets the total number of posts the user has created
func GetTotalUserPosts(db *sql.DB, userID int) (int, error) {
	query := "SELECT COUNT(*) FROM posts WHERE user_id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)

	var totalPosts int
	err = row.Scan(&totalPosts)
	if err != nil {
		return 0, fmt.Errorf("query error: %w", err)
	}

	return totalPosts, nil
}

// GetUserByUsername returns a user from the database by username
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT id, profile_picture, username, age, gender, first_name, last_name, passwrd, email, joined, FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	u := &User{}
	err = row.Scan(&u.ID, &u.ProfilePicture, &u.Username, &u.Age, &u.Gender, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Joined)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user: " + "\"" + username + "\"" + " not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return u, nil
}

// GetUserPosts gets all the posts the user has created
func GetUserPosts(db *sql.DB, userID int) ([]*Post, error) {
	query := `SELECT id, user_id, title, content, category, created
			  FROM posts	
			  WHERE user_id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		p := &Post{}
		err = rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category, &p.Created)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

// GetUserComments gets all the comments the user has created
func GetUserComments(db *sql.DB, userID int) ([]*Comment, error) {
	query := `SELECT id, user_id, post_id, content, created
			  FROM comments
			  WHERE user_id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare statement error: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	comments := []*Comment{}
	for rows.Next() {
		c := &Comment{}
		err = rows.Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &c.Created)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		comments = append(comments, c)
	}

	return comments, nil
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
