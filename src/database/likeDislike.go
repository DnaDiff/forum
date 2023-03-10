package database

import (
	"database/sql"
)

func LikePost(db *sql.DB, userID int, postID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already liked the post
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id=? AND post_id=?", userID, postID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// User has already liked the post, so do nothing
		return nil
	}

	// Insert a new like into the likes table
	stmt, err := tx.Prepare("INSERT INTO likes (user_id, post_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}

	// Update the post's like count
	_, err = tx.Exec("UPDATE posts SET like_count=like_count+1 WHERE id=?", postID)
	if err != nil {
		return err
	}

	// Update the user's like count
	_, err = tx.Exec("UPDATE users SET like_count=like_count+1 WHERE id=?", userID)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
