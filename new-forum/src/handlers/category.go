package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID    string   `json:"ID"`
	Title string   `json:"title"`
	Posts []string `json:"posts"`
}

// Placeholder data
var placeholderCategories = map[string]Category{
	"1": {ID: "1", Title: "Lifestyle", Posts: []string{"123456789", "234567890", "345678901"}},
	"2": {ID: "2", Title: "Entertainment", Posts: []string{"123456789", "234567890", "345678901"}},
	"3": {ID: "3", Title: "Health & Wellness", Posts: []string{"123456789", "234567890", "345678901"}},
	"4": {ID: "4", Title: "Education", Posts: []string{"123456789", "234567890", "345678901"}},
	"5": {ID: "5", Title: "DIY & Crafts", Posts: []string{"123456789", "234567890", "345678901"}},
}

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

func getCategories(db *sql.DB) map[string]Category {
	// Fetch all categories from database below

	// Placeholder
	return placeholderCategories
}

// Get all category postIDs from database
func getCategoriesJSON(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// categories := []Category{}
	// // Convert map to slice (to not get an object)
	// for _, category := range getCategories(db) {
	// 	categories = append(categories, category)
	// }
	// Convert categories to JSON
	categoriesJSON, err := json.Marshal(getCategories(db))
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(categoriesJSON)
}

// Create a new category in database
func createCategory(w http.ResponseWriter, r *http.Request, db *sql.DB, requestData map[string]interface{}) {
	// Expect requestData to contain title
	if len(requestData) == 0 {
		fmt.Println("No data in request body")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get title from requestData
	title, ok := requestData["title"].(string)
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Add post to database below

	// Placeholder
	categoryID := strconv.Itoa(len(placeholderCategories) + 1)
	placeholderCategories[categoryID] = Category{ID: categoryID, Title: strings.ToUpper(title), Posts: []string{}}

	w.WriteHeader(http.StatusCreated)
}

// Delete a category from database
func deleteCategory(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string) {
	// Delete category from database below

	// Placeholder
	delete(placeholderCategories, categoryID)

	w.WriteHeader(http.StatusNoContent)
}

func deletePostFromCategory(db *sql.DB, postID string) {
	// Delete post from category in database below

	// Placeholder
	for _, category := range placeholderCategories {
		for i, post := range category.Posts {
			if post == postID {
				category.Posts = append(category.Posts[:i], category.Posts[i+1:]...)
			}
		}
	}
}
