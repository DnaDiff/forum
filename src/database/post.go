package database

import (
	"database/sql"
	"time"
)

type Post struct {
	ID       int
	UserID   int
	Title    string
	Content  string
	Category string
	Created  time.Time
}

// GetAllPostsByCategory gets all the posts in a category
func GetAllPostsByCategory(db *sql.DB, category string) ([]*Post, error) {
	query := `SELECT id, user_id, title, content, category, created
			  FROM posts	
			  WHERE category = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(category)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		p := &Post{}
		err = rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category, &p.Created)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
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

	return nil
}

// RemovePost removes a post and updates the post count for the user who created the post
func RemovePost(db *sql.DB, postId int) error {
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
	return nil
}
