package handlers

import (
	"database/sql"
	"net/http"
	"time"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)


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

func VerifySession(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

func setSession(token, username string, w http.ResponseWriter) {
	value := map[string]string{
		"token":    token,
		"username": username,
	}
	encoded, err := cookieHandler.Encode("Session", value)
	if err != nil {
		//log.Println("Session encoding failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:     "Session",
		Value:    encoded,
		Path:     "/",
		MaxAge:   sessionExpiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
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

