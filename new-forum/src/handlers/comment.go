package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
POST /api/categories/{categoryId}/posts/{postId}/comments - Add a comment to a specific post from a specific category
DELETE /api/categories/{categoryId}/posts/{postId}/comments/{commentId} - Delete a comment from a specific post in a specific category

Possible parts are [{categoryId}, "posts", {postId}, "comments", {commentId}]
*/

func handleComments(w http.ResponseWriter, r *http.Request, db *sql.DB, parts []string) {
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Println("GET request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments")
		// getComments(w, r, db, parts[0], parts[2])
	case "POST":
		fmt.Println("POST request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments")
		// createComment(w, r, db, requestData, parts[0], parts[2])
	case "DELETE":
		fmt.Println("DELETE request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments/" + parts[4])
		// deleteComment(w, r, db, parts[4])
	}
}
