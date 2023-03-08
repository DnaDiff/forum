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
	Comments  int       `json:"comment_count"`
	Likes     int       `json:"like_count"`
	Dislikes  int       `json:"dislike_count"`
	CreatedAt time.Time `json:"created_at"`
}

// --------------------------------------------Post functions--------------------------------------------

// GetAllPostsByCategory returns all posts in a category
func GetAllPostsByCategory(db *sql.DB, category string) ([]*Post, error) {
	var posts []*Post

	rows, err := db.Query("SELECT id, user_id, title, content, comment_count, like_count, dislike_count, created FROM posts WHERE category = ?", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Comments, &post.Likes, &post.Dislikes, &post.CreatedAt)
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

// CreatePost creates a post and updates the post count for the user who created the post
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

// RemovePost removes a post and all of its comments
func RemovePost(db *sql.DB, postID int) error {

	_, err := db.Exec("DELETE FROM posts WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// if there are any comments on the post, remove them
	_, err1 := db.Exec("DELETE FROM comments WHERE post_id = ?", postID)
	if err1 != nil {
		return err1
	}

	// Update the post count for the user who created the post

	_, err2 := db.Exec("UPDATE users SET post_count = post_count - 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)", postID)
	if err2 != nil {
		return err2
	}

	return nil
}

// LikePost increments the like count for a post
func LikePost(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE posts SET like_count = like_count + 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// Update the like count for the user who created the post
	_, err = db.Exec("UPDATE users SET like_count = like_count + 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)", postID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveLike removes a like from a post
func RemoveLike(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE posts SET like_count = like_count - 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// Update the like count for the user who created the post
	_, err = db.Exec("UPDATE users SET like_count = like_count - 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)", postID)
	if err != nil {
		return err
	}

	return nil
}

// DislikePost increments the dislike count for a post
func DislikePost(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE posts SET dislike_count = dislike_count + 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// Update the dislike count for the user who created the post
	_, err = db.Exec("UPDATE users SET dislike_count = dislike_count + 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)", postID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDislike removes a dislike from a post
func RemoveDislike(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE posts SET dislike_count = dislike_count - 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// Update the dislike count for the user who created the post
	_, err = db.Exec("UPDATE users SET dislike_count = dislike_count - 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)", postID)
	if err != nil {
		return err
	}

	return nil
}
