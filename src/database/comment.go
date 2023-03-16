package database

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID      int
	UserID  int
	PostID  int
	Content string
	Created time.Time
}

// GetComment gets a comment by its ID
func GetComment(db *sql.DB, commentId int) (*Comment, error) {
	query := `SELECT id, user_id, post_id, content, created
			  FROM comments
			  WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	c := &Comment{}
	err = stmt.QueryRow(commentId).Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &c.Created)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CreateComment creates a comment and associates it with a post
func CreateComment(db *sql.DB, comment *Comment) error {
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(comment.UserID, comment.PostID, comment.Content)
	if err != nil {
		return err
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	comment.ID = int(commentId)

	return nil
}

// RemoveComment removes a comment
func RemoveComment(db *sql.DB, commentId int) error {
	stmt, err := db.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentId)
	if err != nil {
		return err
	}

	return nil
}
