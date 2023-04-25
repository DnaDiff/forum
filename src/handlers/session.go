package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	database "github.com/DnaDiff/forum/new-forum/dbfunctions"
)

func HandleCheckSession(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Token string `json:"token"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session, err := database.GetSessionByToken(db, input.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error checking session", http.StatusInternalServerError)
		}
		return
	}
	if time.Now().After(session.ExpiresAt) {
		http.Error(w, "Session expired", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Valid bool `json:"valid"`
	}{
		Valid: true,
	})
}
