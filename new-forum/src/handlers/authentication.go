package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

// CheckCookie checks if the user has a valid cookie
func CheckCookie(w http.ResponseWriter, r *http.Request, db *sql.DB) (int, bool) {
	cookie, err := r.Cookie("user")
	if err != nil {
		fmt.Println("No cookie found")
		return 0, false
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println("Invalid cookie")
		return 0, false
	}

	// if !database.CheckUserID(db, userID) {
	// 	fmt.Println("Invalid cookie")
	// 	return 0, false
	// }

	return userID, true

}

// CreateCookie takes a username and password and creates a cookie for the user
func CreateCookie(w http.ResponseWriter, r *http.Request, db *sql.DB, username string, password string) {

	userID, ok := database.CheckUser(db, username, password)
	fmt.Print(ok)
	if !ok {
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	cookie := http.Cookie{
		Name:     "user",
		Value:    fmt.Sprint(userID),
		Path:     "/",
		MaxAge:   86400, // This cookie will expire after 24 hours
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprint(w, "Cookie for user "+username+" created")
}

func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("Login request received")
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	json.NewDecoder(r.Body).Decode(&requestData)

	// Check if the request data contains the required fields

	if requestData["username"] == nil || requestData["password"] == nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	CreateCookie(w, r, db, requestData["username"].(string), requestData["password"].(string))
}

func HandleRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var requestData map[string]interface{}

	// Decode JSON request body into requestData
	json.NewDecoder(r.Body).Decode(&requestData)

	// Check if the request data contains the required fields

	if requestData["username"] == nil || requestData["password"] == nil || requestData["email"] == nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check if the username is already taken

	// if database.CheckUsername(db, requestData["username"].(string)) {
	// 	http.Error(w, "Username already taken", http.StatusBadRequest)
	// 	return
	// }

	// Check if the email is already taken

	// if database.CheckEmail(db, requestData["email"].(string)) {
	// 	http.Error(w, "Email already taken", http.StatusBadRequest)
	// 	return
	// }

	// Create the user

	var User database.User

	User.Username = requestData["username"].(string)
	User.Password = requestData["password"].(string)
	User.Email = requestData["email"].(string)

	err := database.CreateUser(db, &User)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "User created successfully")
}
