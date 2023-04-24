package database

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	Likes     int       `json:"like_count"`
	Dislikes  int       `json:"dislike_count"`
	CreatedAt time.Time `json:"created_at"`
}

// --------------------------------------------Getting data from comments--------------------------------------------

// GetAllCommentsByPost returns all comments on a post
func GetAllCommentsByPost(db *sql.DB, postID int) ([]*Comment, error) {

	var comments []*Comment

	rows, err := db.Query("SELECT id, user_id, post_id, content, like_count, dislike_count, created FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.Likes, &comment.Dislikes, &comment.CreatedAt)
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

// --------------------------------------------Creating Comments--------------------------------------------

// CreateComment creates a comment on a post and updates the comment count for the user and post
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

// --------------------------------------------Removing Comments--------------------------------------------

// RemoveComment removes a comment
func RemoveComment(db *sql.DB, commentID int) error {

	_, err := db.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		return err
	}

	return nil

}

// --------------------------------------------Like and dislike comment--------------------------------------------

// LikeComment likes a comment
func LikeComment(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE comments SET like_count = like_count + 1 WHERE id = ?", postID)
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

// DislikeComment dislikes a comment
func DislikeComment(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE comments SET dislike_count = dislike_count + 1 WHERE id = ?", postID)
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

// --------------------------------------------Remove like and dislike--------------------------------------------

// RemoveLikeComment removes a like from a comment
func RemoveLikeComment(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE comments SET like_count = like_count - 1 WHERE id = ?", postID)
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

// RemoveDislikeComment removes a dislike from a comment
func RemoveDislikeComment(db *sql.DB, postID int) error {

	_, err := db.Exec("UPDATE comments SET dislike_count = dislike_count - 1 WHERE id = ?", postID)
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
