package database

import (
	"database/sql"
	"log"
	"time"
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
}

func CreateSession(db *sql.DB, session *Session) error {
	result, err := db.Exec("INSERT INTO sessions (user_id, token, expires_at) VALUES (?, ?, ?)", session.UserID, session.Token, session.ExpiresAt)
	if err != nil {
		log.Println("Error inserting session:", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err
	}
	session.ID = int(id)
	log.Printf("Session inserted: %+v\n", session)
	return nil
}

func GetSessionByToken(db *sql.DB, token string) (Session, error) {
	var session Session
	err := db.QueryRow("SELECT id, user_id, token, expires_at FROM sessions WHERE token = ?", token).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt)
	return session, err
}

func ExpireOldSessions(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM sessions WHERE user_id = ? AND expires_at < ?", userID, time.Now())
	return err
}
