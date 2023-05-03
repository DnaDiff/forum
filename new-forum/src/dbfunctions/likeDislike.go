package database

import (
	"database/sql"
	"strconv"
)

// ------------------------------------Get Like/Dislike Post Functions------------------------------------

// Get like count for a post
func GetLikeCount(db *sql.DB, postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id=?", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Get dislike count for a post
func GetDislikeCount(db *sql.DB, postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM dislikes WHERE post_id=?", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAllUsersLikedPost gets all the users that liked the post
func GetAllUsersLikedPost(db *sql.DB, postID int) ([]string, error) {
	rows, err := db.Query("SELECT user_id FROM likes WHERE post_id=?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		users = append(users, strconv.Itoa(userID))
	}

	return users, nil
}

// GetAllUsersDislikedPost gets all the users that disliked the post
func GetAllUsersDislikedPost(db *sql.DB, postID int) ([]string, error) {
	rows, err := db.Query("SELECT user_id FROM dislikes WHERE post_id=?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		users = append(users, strconv.Itoa(userID))
	}

	return users, nil
}

// ------------------------------------Like/Dislike Post Functions------------------------------------

// LikePost adds a like to a post
func LikePost(db *sql.DB, userID int, postID int) error {
	// Remove dislike from post if user has disliked the post
	RemoveDislikePost(db, userID, postID)
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
		// User has already liked the post, remove like and return
		RemoveLikePost(db, userID, postID)
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

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DislikePost adds a dislike to a post and updates the dislike count for the OP
func DislikePost(db *sql.DB, userID int, postID int) error {
	// Remove like from post if user has liked the post
	RemoveLikePost(db, userID, postID)
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
		// User has already disliked the post, remove dislike and return
		RemoveDislikePost(db, userID, postID)
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

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
