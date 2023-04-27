package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

// Handle content requests with handlers from category.go and post.go

func HandleContent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	parts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/"), "/"), "/") // Minimize by trimming empty parts

	fmt.Println("Content request:", r.URL.Path, parts)

	if len(parts) >= 2 && len(parts) <= 3 && parts[1] == "categories" {
		// If the path is /api/categories +{categoryID}, handle with the category handler
		handleCategories(w, r, db, parts[1:]) // parts: ["categories", +{categoryID}]
	} else if len(parts) >= 4 && len(parts) <= 5 && parts[3] == "posts" {
		// If the path is /api/categories/{categoryID}/posts +{postID}, handle with the post handler
		handlePosts(w, r, db, parts[2:]) // parts: [{categoryID}, "posts", +{postID}]
	} else if len(parts) >= 6 && len(parts) <= 7 && parts[5] == "comments" {
		// If the path is /api/categories/{categoryID}/posts/{postID}/comments +{commentID}, handle with the comment handler
		handleComments(w, r, db, parts[2:]) // parts: [{categoryID}, "posts", {postID}, "comments", +{commentID}]
	} else {
		// If the path is invalid or does not start with /api/categories/, return 404
		fmt.Println("Invalid path:", r.URL.Path)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}
