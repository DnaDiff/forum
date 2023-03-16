package database

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID           int
	UserID       int
	PostID       int
	Content      string
	Created      time.Time
	LikeCount    int
	DislikeCount int
}

// CreateComment creates a comment and updates the comment count for the post and user who created the comment
func CreateComment(db *sql.DB, comment *Comment) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	result, err := stmt.Exec(comment.UserID, comment.PostID, comment.Content)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted comment
	commentID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	comment.ID = int(commentID)

	// Update the post's comment count
	stmt, err = db.Prepare("UPDATE posts SET comment_count = comment_count + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.PostID)
	if err != nil {
		return err
	}

	// Update the user's comment count by seeing the amount of comments the user has made
	stmt, err = db.Prepare("UPDATE users SET comment_count = comment_count + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.UserID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveComment removes a comment and updates the comment count for the post and user who created the comment
func RemoveComment(db *sql.DB, commentID int) error {
	// Get the comment's user ID and post ID
	var userID, postID int
	err := db.QueryRow("SELECT user_id, post_id FROM comments WHERE id = ?", commentID).Scan(&userID, &postID)
	if err != nil {
		return err
	}

	// Delete the comment
	stmt, err := db.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID)
	if err != nil {
		return err
	}

	// Update the post's comment count
	stmt, err = db.Prepare("UPDATE posts SET comment_count = comment_count - 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID)
	if err != nil {
		return err
	}

	// Update the user's comment count
	stmt, err = db.Prepare("UPDATE users SET comment_count = comment_count - 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	// Remove any associated likes and dislikes to the comment
	stmt, err = db.Prepare("DELETE FROM likes WHERE comment_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("DELETE FROM dislikes WHERE comment_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID)
	if err != nil {
		return err
	}

	// Update the comment creator's like and dislike count
	var likeCount, dislikeCount int
	err = db.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND comment_id IN (SELECT id FROM comments WHERE user_id = ?)", userID, userID).Scan(&likeCount)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id = ? AND comment_id IN (SELECT id FROM comments WHERE user_id = ?)", userID, userID).Scan(&dislikeCount)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("UPDATE users SET like_count = ?, dislike_count = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeCount, dislikeCount, userID)
	if err != nil {
		return err
	}

	return nil
}
