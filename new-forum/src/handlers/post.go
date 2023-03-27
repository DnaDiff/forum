package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Post struct {
	ID         int    `json:"ID"`
	Category   string `json:"category"`
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

func RetrievePosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "GET" {
		// Fetch all posts from database

		// Placeholder
		posts := []Post{
			{ID: 123456789, Category: "None", Title: "Help me make lasagna", Content: "This is post 123456789", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 123456789, Username: "John_Doe", UserAvatar: DEFAULT_AVATAR},
			{ID: 234567890, Category: "None", Title: "Meditation advice", Content: "This is post 234567890", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 123456789, Username: "John_Doe", UserAvatar: DEFAULT_AVATAR},
			{ID: 345678901, Category: "None", Title: "Party tonight in my discord", Content: "This is post 345678901", Date: "2020-01-01", Comments: []Post{}, Rating: 0, UserID: 345678901, Username: "PARTYBOI", UserAvatar: DEFAULT_AVATAR},
		}

		postsJSON, err := json.Marshal(posts)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(postsJSON)
	} //else if r.Method == "POST" {
	// Create a new post in the database
	//}
}
