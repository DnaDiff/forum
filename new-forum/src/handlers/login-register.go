package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"

	database "github.com/DnaDiff/forum/new-forum/dbfunctions"
)


var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))


func HandleRegister(w http.ResponseWriter, r *http.Request , db *sql.DB) {
	var user database.User
	if r.Body == nil {
		http.Error(w,"Request body is empty", http.StatusBadRequest)
		return
	}
	if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(&user)

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
		log.Printf("User created: %+v\n", user)
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Println("mitä vittua")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
} 


 func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}
	
	if r.Method == "POST" {
		var user database.User
		json.NewDecoder(r.Body).Decode(&user)
		log.Println("Login attempt for user:", user.Username)
		if user.Username == "" && user.Email == "" {
			log.Println("Username or email is required") 
			http.Error(w, "Username or email is required", http.StatusBadRequest)
			return
		}

		var dbUser database.User
		err := db.QueryRow("SELECT id, username, email, passwrd FROM users WHERE username = ? OR email = ?", user.Username, user.Email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Passwrd)
		if err != nil {
			if err == sql.ErrNoRows {
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
			response := map[string]string{"error": "Invalid password"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return			
		}
		log.Println("Password comparison successful") 
		setSession(dbUser.Username, w)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
} 

func HandleLogout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "POST"{
		cookie := &http.Cookie {
			Name: "session",
			Value: "",
			Path: "/",
			MaxAge: -1,
		}
		http.SetCookie(w,cookie)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleVerifySession(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        username, err := getUsernameFromCookie(r)

        if err != nil || username == "" {
            w.WriteHeader(http.StatusUnauthorized)
        } else {
            w.WriteHeader(http.StatusOK)
        }
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func HandleCurrentUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "GET" {
		username, err := getUsernameFromCookie(r)
		if err != nil {
			http.Error(w, "Not logged in", http.StatusUnauthorized)
			return
		}

		response, err := json.Marshal(map[string]string{"username": username})
		if err != nil {
			http.Error(w, "Not logged in", http.StatusUnauthorized)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



func getUsernameFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}

	value := make(map[string]string)
	err = cookieHandler.Decode("session", cookie.Value, &value)
	if err != nil {
		return "", err
	}

	return value["username"], nil
}



func setSession(username string, w http.ResponseWriter) {
	value := map[string]string {
		"username": username,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err != nil {
		log.Println("Session encoding failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name: "session",
		Value: encoded,
		Path: "/",
	}
	http.SetCookie(w, cookie)
}