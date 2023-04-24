package database

import (
	"database/sql"
	"time"
)

type Session struct {
	ID int
	UserID int
	Token string
	ExpiresAt time.Time
}

func (s *Session) CreateSessionTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		token TEXT NOT NULL,
		expires_at TIMESATMP NOT NULL
	)`)
	return err
}

func CreateSession(db *sql.DB, session *Session) error {
	result, err := db.Exec("INSERT INTO sessions (user_id, token, expires_at) (VALUES ? ? ?)",session.UserID, session.Token,session.ExpiresAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	session.ID = int(id)
	return nil
}



func GetSessionByToken(db *sql.DB, token string) (Session, error) {
	var session Session
	err := db.QueryRow("SELECT id, user_id, token, expires_at FROM sessions WHERE token = ?", token).Scan(&session.Token, &session.ExpiresAt)
	return session, err
}

func ExpireOldSessions(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM sessions WHERE user_id = ? AND expires_at < ?", userID, time.Now())
	return err
}

func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, passwrd FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Passwrd)
	return user, err
}