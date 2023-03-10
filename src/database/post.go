package database

import (
	"database/sql"
	"time"
)

type Post struct {
	ID           int
	UserID       int
	Title        string
	Content      string
	Category     string
	Created      time.Time
	LikeCount    int
	DislikeCount int
	CommentCount int
}

// CreatePost creates a post and updates the post count for the user who created the post
func CreatePost(db *sql.DB, post *Post) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO posts(user_id, title, content, category) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	result, err := stmt.Exec(post.UserID, post.Title, post.Content, post.Category)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted post
	postId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.ID = int(postId)

	// Update the user's post count
	stmt, err = db.Prepare("UPDATE users SET post_count = post_count + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.UserID)
	if err != nil {
		return err
	}

	return nil
}

// RemovePost removes a post and updates the post count for the user who created the post
func removePost(db *sql.DB, postId int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Remove all comments associated with the post
	stmt, err := tx.Prepare("DELETE FROM comments WHERE post_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(postId)
	if err != nil {
		return err
	}

	// Remove all likes associated with the post using RemoveLikePost

	// _, err = stmt.Exec(postId)
	// if err != nil {
	//     return err
	// }

	// Remove the post itself
	stmt, err = tx.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(postId)
	if err != nil {
		return err
	}

	// Update the user's post count
	stmt, err = tx.Prepare("UPDATE users SET post_count = post_count - 1 WHERE id = (SELECT user_id FROM posts WHERE id = ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(postId)
	if err != nil {
		return err
	}

	return nil
}
