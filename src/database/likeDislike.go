package database

import (
	"database/sql"
)

// ------------------------------------Like/Dislike Post Functions------------------------------------

// LikePost adds a like to a post and updates the like count for the OP
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

	// Update the post creator's like count
	_, err = tx.Exec("UPDATE users SET like_count=like_count+1 WHERE id=(SELECT user_id FROM posts WHERE id=?)", postID)
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

// DislikePost adds a dislike to a post and updates the dislike count for the OP
func DislikePost(db *sql.DB, userID int, postID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already disliked the post
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id=? AND post_id=? AND comment_id IS NULL", userID, postID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// User has already disliked the post, so do nothing
		return nil
	}

	// Insert a new dislike into the likes table
	stmt, err := tx.Prepare("INSERT INTO dislikes (user_id, post_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}

	// Update the post's dislike count
	_, err = tx.Exec("UPDATE posts SET dislike_count=dislike_count+1 WHERE id=?", postID)
	if err != nil {
		return err
	}

	// Update the post creator's dislike count
	_, err = tx.Exec("UPDATE users SET dislike_count=dislike_count+1 WHERE id=(SELECT user_id FROM posts WHERE id=?)", postID)
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

// RemoveLikePost removes a like from a post and updates the like count for the OP
func RemoveLikePost(db *sql.DB, userID int, postID int) error {
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
	if count == 0 {
		// User has not liked the post, so do nothing
		return nil
	}

	// Remove the like from the likes table
	stmt, err := tx.Prepare("DELETE FROM likes WHERE user_id=? AND post_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}

	// Update the post's like count
	_, err = tx.Exec("UPDATE posts SET like_count=like_count-1 WHERE id=?", postID)
	if err != nil {
		return err
	}

	// Update the post creator's like count
	_, err = tx.Exec("UPDATE users SET like_count=like_count-1 WHERE id=(SELECT user_id FROM posts WHERE id=?)", postID)
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

// RemoveDisLikePost removes a dislike from a post and updates the dislike count for the OP
func RemoveDislikePost(db *sql.DB, userID int, postID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already disliked the post
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id=? AND post_id=?", userID, postID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// User has not disliked the post, so do nothing
		return nil
	}

	// Remove the dislike from the dislikes table
	stmt, err := tx.Prepare("DELETE FROM dislikes WHERE user_id=? AND post_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}

	// Update the post's dislike count
	_, err = tx.Exec("UPDATE posts SET dislike_count=dislike_count-1 WHERE id=?", postID)
	if err != nil {
		return err
	}

	// Update the post creator's dislike count
	_, err = tx.Exec("UPDATE users SET dislike_count=dislike_count-1 WHERE id=(SELECT user_id FROM posts WHERE id=?)", postID)
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

// ------------------------------------Like/Dislike Comment Functions------------------------------------

// LikeComment adds a like to a comment and updates the like count for the OP
func LikeComment(db *sql.DB, userID int, commentID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already liked the comment
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id=? AND comment_id=?", userID, commentID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// User has already liked the comment, so do nothing
		return nil
	}

	// Insert a new like into the likes table
	stmt, err := tx.Prepare("INSERT INTO likes (user_id, comment_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, commentID)
	if err != nil {
		return err
	}

	// Update the comment's like count
	_, err = tx.Exec("UPDATE comments SET like_count=like_count+1 WHERE id=?", commentID)
	if err != nil {
		return err
	}

	// Update the comment creator's like count
	_, err = tx.Exec("UPDATE users SET like_count=like_count+1 WHERE id=(SELECT user_id FROM comments WHERE id=?)", commentID)
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

// DislikeComment adds a dislike to a comment and updates the dislike count for the OP
func DislikeComment(db *sql.DB, userID int, commentID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// Check if the user has already disliked the comment
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id=? AND comment_id=?", userID, commentID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// User has already disliked the comment, so do nothing
		return nil
	}

	// Insert a new dislike into the dislikes table
	stmt, err := tx.Prepare("INSERT INTO dislikes (user_id, comment_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, commentID)
	if err != nil {
		return err
	}

	// Update the comment's dislike count
	_, err = tx.Exec("UPDATE comments SET dislike_count=dislike_count+1 WHERE id=?", commentID)
	if err != nil {
		return err
	}

	// Update the comment creator's dislike count
	_, err = tx.Exec("UPDATE users SET dislike_count=dislike_count+1 WHERE id=(SELECT user_id FROM comments WHERE id=?)", commentID)
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

// RemoveLikeComment removes a like from a comment and updates the like count for the OP
func RemoveLikeComment(db *sql.DB, userID int, commentID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already liked the comment
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id=? AND comment_id=?", userID, commentID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// User has not liked the comment, so do nothing
		return nil
	}

	// Remove the like from the likes table
	stmt, err := tx.Prepare("DELETE FROM likes WHERE user_id=? AND comment_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, commentID)
	if err != nil {
		return err
	}

	// Update the comment's like count
	_, err = tx.Exec("UPDATE comments SET like_count=like_count-1 WHERE id=?", commentID)
	if err != nil {
		return err
	}

	// Update the comment creator's like count
	_, err = tx.Exec("UPDATE users SET like_count=like_count-1 WHERE id=(SELECT user_id FROM comments WHERE id=?)", commentID)
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

// RemoveDislikeComment removes a dislike from a comment and updates the dislike count for the OP
func RemoveDislikeComment(db *sql.DB, userID int, commentID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if the user has already disliked the comment
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id=? AND comment_id=?", userID, commentID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// User has not disliked the comment, so do nothing
		return nil
	}

	// Remove the dislike from the dislikes table
	stmt, err := tx.Prepare("DELETE FROM dislikes WHERE user_id=? AND comment_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, commentID)
	if err != nil {
		return err
	}

	// Update the comment's dislike count
	_, err = tx.Exec("UPDATE comments SET dislike_count=dislike_count-1 WHERE id=?", commentID)
	if err != nil {
		return err
	}

	// Update the comment creator's dislike count
	_, err = tx.Exec("UPDATE users SET dislike_count=dislike_count-1 WHERE id=(SELECT user_id FROM comments WHERE id=?)", commentID)
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
