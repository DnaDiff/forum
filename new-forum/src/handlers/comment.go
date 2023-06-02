package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

/*
POST /api/categories/{categoryId}/posts/{postId}/comments - Add a comment to a specific post from a specific category
DELETE /api/categories/{categoryId}/posts/{postId}/comments/{commentId} - Delete a comment from a specific post in a specific category

Possible parts are [{categoryId}, "posts", {postId}, "comments", {commentId}]
*/

func handleComments(w http.ResponseWriter, r *http.Request, db *sql.DB, parts []string) {

	var requestData map[string]any

	userID, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to create a post", http.StatusUnauthorized)
		return
	}

	// Decode JSON request body into requestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	postID, _ := strconv.Atoi(parts[2])

	insertComment := database.CommentInsertDB{
		UserID:  userID,
		PostID:  postID,
		Content: requestData["content"].(string),
	}

	responseSuccessComment := database.CommentDB{
		ID:      -1,
		UserID:  userID,
		PostID:  postID,
		Content: requestData["content"].(string),
		Created: time.Now(),
	}

	switch r.Method {
	case "GET":
		log.Println("GET request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments")
		// getComments(w, r, db, parts[0], parts[2])
	case "POST":
		log.Println("POST request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments")
		responseSuccessComment.ID, err = database.CreateComment(db, &insertComment)
		if err != nil {
			http.Error(w, err.Error()+"Comment could not be created", http.StatusInternalServerError)
			return
		}
		// respond with json and status created
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responseSuccessComment)
	case "DELETE":
		log.Println("DELETE request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments/" + parts[4])
		// deleteComment(w, r, db, parts[4])
	}
}
