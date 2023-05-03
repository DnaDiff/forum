package database

import (
	"database/sql"
	"fmt"
)

type Category struct {
	ID    int
	Title string
}

func GetCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT id, title FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func AddCategory(db *sql.DB, title string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if category already exists
	var existingID int
	err = tx.QueryRow("SELECT id FROM categories WHERE title = ?", title).Scan(&existingID)
	if err == nil {
		// Category already exists, return error
		return fmt.Errorf("category with title '%s' already exists", title)
	} else if err != sql.ErrNoRows {
		// Unexpected error occurred
		return err
	}

	// Insert new category
	res, err := tx.Exec("INSERT INTO categories (title) VALUES (?)", title)
	if err != nil {
		return err
	}
	categoryID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Added category with ID %d and title '%s'\n", categoryID, title)
	return nil
}

func RemoveCategory(db *sql.DB, categoryID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if category exists
	var existingTitle string
	err = tx.QueryRow("SELECT title FROM categories WHERE id = ?", categoryID).Scan(&existingTitle)
	if err == sql.ErrNoRows {
		// Category does not exist, return error
		return fmt.Errorf("category with ID %d does not exist", categoryID)
	} else if err != nil {
		// Unexpected error occurred
		return err
	}

	// Delete category
	_, err = tx.Exec("DELETE FROM categories WHERE id = ?", categoryID)
	if err != nil {
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Removed category with ID %d and title '%s'\n", categoryID, existingTitle)
	return nil
}

