package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

const sessionExpiration = 3600


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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	err = database.CreateUser(db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new session for the registered user
	sessionToken, err := createNewSession(db, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSession(sessionToken, user.Username, w)
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

	dbUser, err := database.GetUserByUsername(db, user.Username)
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
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		log.Println("Password comparison failed:", err)
		response := map[string]string{"Error": "Invalid password"}
		w.WriteHeader(http.StatusUnauthorized)
		//----

		json.NewEncoder(w).Encode(response)
		return
	}
	sessionToken, err := createNewSession(db, dbUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Setting up the session!")
		setSession(sessionToken, dbUser.Username, w)
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"Success": "Logged in"}
		json.NewEncoder(w).Encode(response)

}

func HandleLogout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	clearSession(w)
	w.WriteHeader(http.StatusOK)
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


func GenerateRandomToken() string {
	return base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32))
}
