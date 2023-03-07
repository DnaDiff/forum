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

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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

// --------------------------------------------Post functions--------------------------------------------

func GetAllPosts(db *sql.DB) ([]*Post, error) {
	var posts []*Post

	rows, err := db.Query("SELECT id, user_id, title, content, created FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(db *sql.DB, userID int, title string, content string) error {
	_, err := db.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", userID, title, content)
	if err != nil {
		return err
	}

	// Update the post count for the user who created the post
	_, err = db.Exec("UPDATE users SET post_count = post_count + 1 WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}

func GetPostByID(db *sql.DB, postID int) (*Post, error) {
	var post Post

	err := db.QueryRow("SELECT id, user_id, title, content, created_at FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func CreateComment(db *sql.DB, userID int, postID int, content string) error {
	_, err := db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)", userID, postID, content)
	if err != nil {
		return err
	}

	// Update the comment count for the user who created the comment
	_, err = db.Exec("UPDATE users SET comment_count = comment_count + 1 WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}

func GetCommentByID(db *sql.DB, commentID int) (*Comment, error) {
	var comment Comment

	err := db.QueryRow("SELECT id, user_id, post_id, content, created_at FROM comments WHERE id = ?", commentID).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
