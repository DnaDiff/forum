package database

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Comments  int       `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// --------------------------------------------Post functions--------------------------------------------

func GetAllPostsByCategory(db *sql.DB, category string) ([]*Post, error) {
	var posts []*Post

	rows, err := db.Query("SELECT id, user_id, title, content, created FROM posts WHERE category = ?", category)
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

func GetAllCommentsByPostID(db *sql.DB, postID int) ([]*Comment, error) {
	var comments []*Comment

	rows, err := db.Query("SELECT id, user_id, post_id, content, created FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

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

func CreatePost(db *sql.DB, userID int, title string, content string, category string) error {
	_, err := db.Exec("INSERT INTO posts (user_id, title, content, category) VALUES (?, ?, ?, ?)", userID, title, content, category)
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

	// Update the comment count for the post that the comment was created on
	_, err = db.Exec("UPDATE posts SET comment_count = comment_count + 1 WHERE id = ?", postID)
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
