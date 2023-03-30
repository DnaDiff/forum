package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Post struct {
	ID         int    `json:"ID"`
	ParentID   string `json:"parentID"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Date       string `json:"date"`
	Comments   []Post `json:"comments"`
	Rating     int    `json:"rating"`
	UserID     int    `json:"userID"`
	Username   string `json:"username"`
	UserAvatar string `json:"userAvatar"`
}

const DEFAULT_AVATAR = "https://st3.depositphotos.com/6672868/13701/v/600/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"

func HandlePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var postID, postAction string
	parts := strings.Split(r.URL.Path, "/")

	// Assign postID and postAction whether the request includes a specific post ID or action
	if len(parts) >= 4 {
		postID = strings.Split(r.URL.Path, "/")[3] // /api/posts/{id}
		if len(parts) == 5 {
			postAction = strings.Split(r.URL.Path, "/")[4] // /api/posts/{id}/{action}
		}
	}

	switch r.Method {
	case "GET":
		if postID != "" {
			fmt.Println("GET request to /api/posts/" + postID)
			RetrievePost(w, r, db)
		} else {
			fmt.Println("GET request to /api/posts")
			RetrievePosts(w, r, db)
		}
	case "POST":
		if postAction == "create" {
			fmt.Printf("POST request to /api/posts/%s\n", postAction)
			CreatePost(w, r, db)
		} else if postAction == "comment" {
			fmt.Printf("POST request to /api/posts/%s/%s\n", postID, postAction)
			CommentPost(w, r, db)
		} else if postAction == "upvote" {
			fmt.Printf("POST request to /api/posts/%s/%s\n", postID, postAction)
			UpvotePost(w, r, db)
		}
	case "PUT":
		if postAction == "downvote" {
			fmt.Printf("PUT request to /api/posts/%s/%s\n", postID, postAction)
			DownvotePost(w, r, db)
		}
	case "DELETE":
		if postAction == "delete" {
			fmt.Printf("DELETE request to /api/posts/%s/%s\n", postID, postAction)
			DeletePost(w, r, db)
		}
	}
}

// Retrieve all posts and their data from database
func RetrievePosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Fetch all posts from database

	// Placeholder
	posts := []Post{
		{ID: 123456789, ParentID: "None", Title: "Help me make lasagna", Content: "This is post 123456789", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 123456789, Username: "John_Doe", UserAvatar: DEFAULT_AVATAR},
		{ID: 234567890, ParentID: "None", Title: "Meditation advice", Content: "This is post 234567890", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 123456789, Username: "John_Doe", UserAvatar: DEFAULT_AVATAR},
		{ID: 345678901, ParentID: "None", Title: "Party tonight in my discord", Content: "This is post 345678901", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 345678901, Username: "PARTYBOI", UserAvatar: DEFAULT_AVATAR},
	}

	postsJSON, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(postsJSON)
}

// Retrieve a specific post and its data from database
func RetrievePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Fetch post from database
	// fmt.Println("GET request to /api/posts/{id}")

	// Placeholder
	if r.URL.Path == "/api/posts/123456789" {
		post := Post{ID: 123456789, ParentID: "None", Title: "Help me make test", Content: "This is post 123456789", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 123456789, Username: "John_Doe", UserAvatar: DEFAULT_AVATAR}

		postJSON, err := json.Marshal(post)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(postJSON)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Add post to database
	// fmt.Println("POST request to /api/posts/create")

	// Placeholder
	w.WriteHeader(http.StatusOK)
}

func CommentPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Add new post to database with parentID and add ID to parent post's comments
	// fmt.Println("POST request to /api/posts/{id}/comment")

	// Placeholder
	w.WriteHeader(http.StatusOK)
}

func UpvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Add upvote from user to database
	// fmt.Println("POST request to /api/posts/{id}/upvote")

	// Placeholder
	w.WriteHeader(http.StatusOK)
}

func DownvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Remove upvote from user in database
	// fmt.Println("PUT request to /api/posts/{id}/downvote")

	// Placeholder
	w.WriteHeader(http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Delete post from database
	// fmt.Println("DELETE request to /api/posts/{id}/delete")

	// Placeholder
	w.WriteHeader(http.StatusOK)
}
