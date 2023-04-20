package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"database/sql"

	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"

	database "github.com/DnaDiff/forum/new-forum/dbfunctions"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

const sessionExpiration = 3600


func createNewSession(db *sql.DB, userId int) (string, error) {
	sessionToken := GenerateRandomToken()

	expiresAt := time.Now().Add(time.Duration(sessionExpiration) * time.Second)

	err := database.ExpireOldSessions(db, userId)
	if err != nil {
		return "", err
	}
	session := &database.Session{
		UserID: userId,
		Token: sessionToken,
		ExpiresAt: expiresAt,
	}
	err = database.CreateSession(db, session)
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}


func HandleRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Passwrd), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Passwrd = string(hashedPassword)

	err = database.CreateUser(db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
		setSession(user.Username, w)
		w.WriteHeader(http.StatusCreated)
	
}

func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Decoded user:", user)

	log.Println("Login attempt for user:", user.Username)
	if user.Username == "" && user.Email == "" {
		log.Println("Username or email is required")
		http.Error(w, "Username or email is required", http.StatusBadRequest)
		return
	}

	dbUser, err := getUserFromDatabase(db, user.Username, user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			log.Println("Query error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	log.Printf("User found in the database: %+v\n", dbUser)
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Passwrd), []byte(user.Passwrd))
	if err != nil {
		log.Println("Password comparison failed:", err)
		response := map[string]string{"Error": "Invalid password"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	return
	}
	sessionToken, err := createNewSession(db, dbUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Password comparison successful")
	setSession(sessionToken, w)
	w.WriteHeader(http.StatusOK)

}

func HandleLogout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	clearSession(w)
	w.WriteHeader(http.StatusOK)
}



func HandleVerifySession(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sessionToken, err := getSessionTokenFromCookie(r)
	if err != nil || sessionToken == ""{
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func HandleCurrentUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
	}
	sessionToken, err := getSessionTokenFromCookie(r)
	if err != nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}
	session, err := database.GetSessionByToken(db, sessionToken)
	if err != nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}
	dbUser, err := database.GetUserByID(db, session.UserID)
	if err != nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}
	response, err := json.Marshal(map[string]string{"username": dbUser.Username})
	if err != nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(response)
}


func getUserFromDatabase(db *sql.DB, username, email string) (database.User, error) {
	var dbUser database.User
	err := db.QueryRow("SELECT id, username, email, passwrd FROM users WHERE username = ? OR email = ?", username, email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Passwrd)
	return dbUser, err
}

func getUsernameFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("Session")
	if err != nil {
	return "", err
	}
	value := make(map[string]string)
	err = cookieHandler.Decode("Session", cookie.Value, &value)
	if err != nil {
	return "", err
	}

	return value["username"], nil
}

func setSession(token string, w http.ResponseWriter) {
	value := map[string]string {
		"token": token,
	}
	encoded, err := cookieHandler.Encode("Session", value)
	if err != nil {
		log.Println("Session encoding failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie {
		Name: "Session",
		Value: encoded,
		Path: "/",
		MaxAge: sessionExpiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func getSessionTokenFromCookie(r *http.Request) (string,error) {
	cookie, err := r.Cookie("Session")
	if err != nil {
		return "", err
	}
	value := make(map[string]string)
	err = cookieHandler.Decode("Session", cookie.Value, &value)
	if err != nil {
		return "", err
	}
	return value["token"], nil
}

func clearSession(w http.ResponseWriter){
	cookie := &http.Cookie {
		Name: "Session",
		Value: "",
		Path: "/",
		MaxAge: -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func GenerateRandomToken() string {
	return base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32))
}

