package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

type Rating struct {
	ID     string `json:"id"`
	UserID string `json:"userID"`
	PostID string `json:"postID"`
	Action string `json:"action"`
}

func GetUpvotes(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) []string {
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil
	}
	upvotes, err := database.GetAllUsersLikedPost(db, postIDInt)
	if err != nil {
		fmt.Printf("Error getting all users who upvoted post: %v\n", err)
		http.Error(w, "Failed to get all users who upvoted the post", http.StatusInternalServerError)
		return nil
	}

	return upvotes
}

func GetDownvotes(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) []string {
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil
	}
	downvotes, err := database.GetAllUsersDislikedPost(db, postIDInt)
	if err != nil {
		fmt.Printf("Error getting all users who downvoted post: %v\n", err)
		http.Error(w, "Failed to get all users who downvoted the post", http.StatusInternalServerError)
		return nil
	}
	return downvotes
}
