package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

type Post struct {
	ID         string   `json:"ID"`
	CategoryID string   `json:"categoryID"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Created    string   `json:"created"`
	Rating     int      `json:"rating"`
	Upvotes    []string `json:"upvotes"`
	Downvotes  []string `json:"downvotes"`
	UserID     string   `json:"userID"`
	Username   string   `json:"username"`
	UserAvatar string   `json:"userAvatar"`
}

func handlePosts(w http.ResponseWriter, r *http.Request, db *sql.DB, parts []string) {
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	json.NewDecoder(r.Body).Decode(&requestData)

	switch r.Method {
	case "GET":
		if len(parts) == 2 {
			log.Println("GET request to /api/categories/" + parts[0] + "/posts")
			getPostsJSON(w, r, db, parts[0])
		}
	case "POST":
		if len(parts) == 2 {
			log.Println("POST request to /api/categories/" + parts[0] + "/posts")
			createPost(w, r, db, parts[0], requestData)
		} else if len(parts) == 3 {
			log.Println("POST request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			upvotePost(w, r, db, parts[2])
		} else if len(parts) == 4 && parts[3] == "comments" {
			log.Println("POST request to /api/categories/" + parts[0] + "/posts/" + parts[2] + "/comments")
			commentPost(w, r, db, parts[2], requestData)
		}
	case "DELETE":
		if len(parts) == 3 {
			log.Println("DELETE request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			deletePost(w, r, db, parts[2])
		}
	case "PUT":
		if len(parts) == 3 {
			log.Println("PUT request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			downvotePost(w, r, db, parts[2])
		}
	}
}

// Get requested posts and their data from database
func getPostsJSON(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string) {
	var categoryIDInt, err = strconv.Atoi(categoryID)
	if err != nil {
		fmt.Printf("Invalid category ID: %v\n", err)
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	// Get all IDs in category from database
	postIDs, err := database.GetAllPostIDsByCategoryID(db, categoryIDInt)
	if err != nil {
		fmt.Printf("Error getting post IDs: %v\n", err)
		http.Error(w, "Error getting post IDs", http.StatusInternalServerError)
		return
	}
	// Fetch all posts from database

	var posts []Post
	for _, postID := range postIDs {
		newPost := getPost(w, r, db, postID)
		if newPost.ID == "" {
			continue
		}
		posts = append(posts, getPost(w, r, db, postID))
	}
	postsJSON, err := json.Marshal(posts)
	if err != nil {
		fmt.Printf("Error marshalling posts: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(postsJSON)
}

// Get a specific post and its updated data from database
func getPost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) Post {
	// Fetch post from database
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return Post{}
	}

	postDB, err := database.GetPost(db, postIDInt)
	if err != nil {
		fmt.Printf("Post["+postID+"] not found: %v\n", err)
		http.Error(w, "Post["+postID+"] not found", http.StatusNotFound)
		return Post{}
	}

	// Convert postDB to post
	post := Post{
		ID:         strconv.Itoa(postDB.ID),
		CategoryID: strconv.Itoa(postDB.CategoryID),
		Title:      postDB.Title,
		Content:    postDB.Content,
		Created:    postDB.Created.Format("2006-01-02 15:04:05"),
		UserID:     strconv.Itoa(postDB.UserID),
	}

	var user User = getUser(w, r, db, post.UserID)
	post.Username = strings.ToUpper(user.Username)
	post.UserAvatar = user.Avatar

	post.Upvotes = GetUpvotes(w, r, db, postID)
	post.Downvotes = GetDownvotes(w, r, db, postID)
	post.Rating = len(post.Upvotes) - len(post.Downvotes)

	return post
}

// Create a new post under a specified category in database
func createPost(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string, requestData map[string]interface{}) {
	//Check if user is logged in
	userID, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to create a post", http.StatusUnauthorized)
		return
	}
	// Expect requestData to contain data for post
	if len(requestData) == 0 {
		fmt.Println("No data in request body")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get title, content and userID from requestData
	title, ok := requestData["title"].(string)
	if !ok {
		fmt.Println("No title found in request body")
		http.Error(w, "No title found", http.StatusBadRequest)
		return
	}

	content, ok := requestData["content"].(string)
	if !ok {
		log.Println("No content found in request body")
		http.Error(w, "No content found", http.StatusBadRequest)
		return
	}

	// !!! When AUTHENTICATION is implemented, get userID from token (in request header) !!!

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		fmt.Printf("Error converting categoryID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Add post to database
	err = database.CreatePost(db, &database.PostDB{UserID: userID, Title: title, Content: content, CategoryID: categoryIDInt})
	if err != nil {
		fmt.Printf("Error creating post: %v\n", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Post created"}`))
}

// Delete a post from a specified category in database
func deletePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	//Check if user is logged in
	_, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to delete a post", http.StatusUnauthorized)
		return
	}

	// Remove post from database below [REQUIRE OWNER AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = database.RemovePost(db, postIDInt)
	if err != nil {
		fmt.Printf("Error removing post[%v]: %v\n", postID, err)
		http.Error(w, "Failed to remove post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func commentPost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string, requestData map[string]interface{}) {
	log.Println("I was here like a good boy")
	//Check if user is logged in
	userID, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to comment on a post", http.StatusUnauthorized)
		return
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Add new post to database and add ID to parent post's comments [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]

	content, ok := requestData["content"].(string)
	if !ok {
		fmt.Println("No content found in request body")
		http.Error(w, "No content found", http.StatusBadRequest)
		return
	}

	commentID, err := database.CreateComment(db, &database.CommentInsertDB{UserID: userID, PostID: postIDInt, Content: content})
	if err != nil {
		fmt.Printf("Error creating comment: %v\n", err)
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	comment := database.CommentDB{
		ID:      commentID,
		UserID:  userID,
		PostID:  postIDInt,
		Content: content,
		Created: time.Now(),
	}

	response, err := json.Marshal(comment)
	if err != nil {
		fmt.Printf("Error marshalling comment: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println(response)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func upvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	//Check if user is logged in
	userID, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to upvote a post", http.StatusUnauthorized)
		return
	}
	// Add upvote from user to database [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = database.LikePost(db, userID, postIDInt) // userID = 0 until authentication is implemented
	if err != nil {
		fmt.Printf("Error upvoting post[%v]: %v\n", postID, err)
		http.Error(w, "Failed to upvote post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func downvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	//Check if user is logged in
	userID, loggedIn := CheckCookie(w, r, db)
	if !loggedIn {
		http.Error(w, "You must be logged in to downvote a post", http.StatusUnauthorized)
		return
	}
	// Remove upvote from user in database [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Printf("Error converting postID to int: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = database.DislikePost(db, userID, postIDInt) // userID = 0 until authentication is implemented
	if err != nil {
		fmt.Printf("Error downvoting post[%v]: %v\n", postID, err)
		http.Error(w, "Failed to downvote post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
