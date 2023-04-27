package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Post struct {
	ID         string   `json:"ID"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Date       string   `json:"date"`
	Comments   []string `json:"comments"`
	Rating     int      `json:"rating"`
	UserID     string   `json:"userID"`
	Username   string   `json:"username"`
	UserAvatar string   `json:"userAvatar"`
}

var placeholderPosts = map[string]Post{
	"123456789": {ID: "123456789", Title: "Help me make lasagna", Content: "This is post 123456789", Date: "2020-01-01", Comments: []string{}, Rating: 0, UserID: "123456789", Username: "", UserAvatar: ""},
	"234567890": {ID: "234567890", Title: "Meditation advice", Content: "This is post 234567890", Date: "2020-01-01", Comments: []string{}, Rating: 0, UserID: "123456789", Username: "", UserAvatar: ""},
	"345678901": {ID: "345678901", Title: "Party tonight in my discord", Content: "This is post 345678901", Date: "2020-01-01", Comments: []string{}, Rating: 0, UserID: "345678901", Username: "", UserAvatar: ""},
}

/*
GET /api/categories/{categoryId}/posts - View posts from a specific category
POST /api/categories/{categoryId}/posts - Create a post in a specific category
DELETE /api/categories/{categoryId}/posts/{postId} - Delete a post from a specific category
POST /api/categories/{categoryId}/posts/{postId} [requestData "action": "upvote"] - Upvote a specific post from a specific category
PUT /api/categories/{categoryId}/posts/{postId} [requestData "action": "downvote"] - Downvote a specific post from a specific category

Possible parts are [{categoryId}, "posts", {postId}]
*/

func handlePosts(w http.ResponseWriter, r *http.Request, db *sql.DB, parts []string) {
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	json.NewDecoder(r.Body).Decode(&requestData)

	switch r.Method {
	case "GET":
		if len(parts) == 2 {
			fmt.Println("GET request to /api/categories/" + parts[0] + "/posts")
			getPostsJSON(w, r, db, parts[0], requestData)
		}
	case "POST":
		if len(parts) == 2 {
			fmt.Println("POST request to /api/categories/" + parts[0] + "/posts")
			createPost(w, r, db, parts[0], requestData)
		} else if len(parts) == 3 {
			fmt.Println("POST request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			upvotePost(w, r, db, parts[2])
		}
	case "DELETE":
		if len(parts) == 3 {
			fmt.Println("DELETE request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			deletePost(w, r, db, parts[2])
		}
	case "PUT":
		if len(parts) == 3 {
			fmt.Println("PUT request to /api/categories/" + parts[0] + "/posts/" + parts[2])
			downvotePost(w, r, db, parts[2])
		}
	}

}

// Get requested posts and their data from database
func getPostsJSON(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string, requestData map[string]interface{}) {
	var postIDs []interface{}
	var ok bool

	// If requestData is empty, return all posts in category
	if len(requestData) == 0 {
		// Get all IDs in category and convert to []interface{}
		for _, postID := range getCategories(db)[categoryID].Posts {
			postIDs = append(postIDs, postID)
		}
	} else {
		// Get postIDs from requestData
		postIDs, ok = requestData["postIDs"].([]interface{})
		if !ok {
			http.Error(w, "postIDs not found", http.StatusBadRequest)
			return
		}
	}

	// Fetch all posts from database

	// Placeholder
	var posts []Post
	for _, postID := range postIDs {
		newPost := getPost(w, r, db, postID.(string))
		if newPost.ID == "" {
			continue
		}
		posts = append(posts, getPost(w, r, db, postID.(string)))
	}

	postsJSON, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(postsJSON)
}

// Get a specific post and its updated data from database
func getPost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) Post {
	// Fetch post from database

	// Placeholder
	var post, ok = placeholderPosts[postID]
	if !ok {
		http.Error(w, "Post["+postID+"] not found", http.StatusNotFound)
		return Post{}
	}

	var user User = getUser(w, r, db, post.UserID)
	post.Username = strings.ToUpper(user.Username)
	post.UserAvatar = user.Avatar

	return post
}

// Create a new post under a specified category in database
func createPost(w http.ResponseWriter, r *http.Request, db *sql.DB, categoryID string, requestData map[string]interface{}) {
	// Expect requestData to contain data for post
	if len(requestData) == 0 {
		fmt.Println("No data in request body")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get title, content and userID from requestData
	title, ok := requestData["title"].(string)
	if !ok {
		http.Error(w, "No title found", http.StatusBadRequest)
		return
	}

	content, ok := requestData["content"].(string)
	if !ok {
		http.Error(w, "No content found", http.StatusBadRequest)
		return
	}

	// !!! When AUTHENTICATION is implemented, get userID from token (in request header) instead !!!
	userID, ok := requestData["userID"].(string)
	if !ok {
		http.Error(w, "No userID found", http.StatusBadRequest)
		return
	}
	// Add post to database below

	// Placeholder
	postID := strconv.Itoa(len(placeholderPosts) + 1)
	placeholderPosts[postID] = Post{ID: postID, Title: strings.ToUpper(title), Content: content, Date: time.Now().Local().String(), Comments: []string{}, Rating: 0, UserID: userID, Username: "", UserAvatar: ""} // Username and UserAvatar is set when post is fetched from database

	category := placeholderCategories[categoryID]
	category.Posts = append(category.Posts, postID)
	placeholderCategories[categoryID] = category

	w.WriteHeader(http.StatusCreated)
}

// Delete a post from a specified category in database
func deletePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	// Remove post from database below [REQUIRE OWNER AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]
	deletePostFromCategory(db, postID)

	// Placeholder
	delete(placeholderPosts, postID)
	w.WriteHeader(http.StatusOK)
}

func commentPost(w http.ResponseWriter, r *http.Request, db *sql.DB, parentID string) {
	// Add new post to database and add ID to parent post's comments [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]

	// Placeholder
	w.WriteHeader(http.StatusOK)
}

func upvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	// Add upvote from user to database [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]

	// Placeholder
	post := placeholderPosts[postID]
	post.Rating++
	placeholderPosts[postID] = post
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"rating": ` + strconv.Itoa(placeholderPosts[postID].Rating) + `}`))
}

func downvotePost(w http.ResponseWriter, r *http.Request, db *sql.DB, postID string) {
	// Remove upvote from user in database [REQUIRE AUTHENTICATION THROUGH TOKEN FROM REQUEST HEADER]

	// Placeholder
	post := placeholderPosts[postID]
	post.Rating--
	placeholderPosts[postID] = post
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"rating": ` + strconv.Itoa(placeholderPosts[postID].Rating) + `}`))
}
