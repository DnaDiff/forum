package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

type Category struct {
	ID    string `json:"ID"`
	Title string `json:"title"`
}

// Placeholder data
// var placeholderCategories = map[string]Category{
// 	"1": {ID: "1", Title: "Lifestyle", Posts: []string{"123456789", "234567890", "345678901"}},
// 	"2": {ID: "2", Title: "Entertainment", Posts: []string{"123456789", "234567890", "345678901"}},
// 	"3": {ID: "3", Title: "Health & Wellness", Posts: []string{"123456789", "234567890", "345678901"}},
// 	"4": {ID: "4", Title: "Education", Posts: []string{"123456789", "234567890", "345678901"}},
// 	"5": {ID: "5", Title: "DIY & Crafts", Posts: []string{"123456789", "234567890", "345678901"}},
// }

/*
POST /api/categories - Create a category
DELETE /api/categories/{categoryId} - Delete a category
GET /api/categories - List categories

Possible parts are ["categories", {categoryId}]
*/
func handleCategories(w http.ResponseWriter, r *http.Request, db *sql.DB, parts []string) {
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	json.NewDecoder(r.Body).Decode(&requestData)

	switch r.Method {
	case "GET":
		if len(parts) == 1 && parts[0] == "categories" {
			fmt.Println("GET request to /api/categories")
			getCategoriesJSON(w, r, db)
		} else {
			fmt.Println("Invalid GET request to " + r.URL.Path)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	case "POST":
		fmt.Println("POST request to /api/categories")
		createCategory(w, r, db, requestData)
	case "DELETE":
		fmt.Println("DELETE request to /api/categories/" + parts[1])
		deleteCategory(w, r, db, parts[1])
	}
}

func getCategories(db *sql.DB) []Category {
	// Fetch all categories from database below
	categoriesDB, err := database.GetCategories(db)
	if err != nil {
		fmt.Printf("Error getting categories from database: %v\n", err)
		return nil
	}

	// Convert categoriesDB to categories
	categories := []Category{}
	for _, categoryDB := range categoriesDB {
		categories = append(categories, Category{ID: strconv.Itoa(categoryDB.ID), Title: categoryDB.Title})
	}

	return categories
}

// Get all category postIDs from database
func getCategoriesJSON(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Convert categories to JSON
	categoriesJSON, err := json.Marshal(getCategories(db))
	if err != nil {
		fmt.Printf("Error marshalling categories to JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(categoriesJSON)
}

// Create a new category in database
func createCategory(w http.ResponseWriter, r *http.Request, db *sql.DB, requestData map[string]interface{}) {
	// Expect requestData to contain title
	if len(requestData) == 0 {
		fmt.Printf("No data in request body\n")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get title from requestData
	title, ok := requestData["title"].(string)
	if !ok {
		fmt.Printf("Error getting title from request body\n")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Add category to database
	err := database.AddCategory(db, strings.ToUpper(title))
	if err != nil {
		fmt.Printf("Error adding category to database: %v\n", err)
		http.Error(w, "Error adding category to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete a category from database
func deleteCategory(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string) {
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		fmt.Printf("Error converting categoryID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Delete category from database below
	err = database.RemoveCategory(db, categoryIDInt)
	if err != nil {
		fmt.Printf("Error removing category from database: %v\n", err)
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
