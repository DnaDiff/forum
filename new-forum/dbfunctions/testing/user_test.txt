package database

import (
	"database/sql"
	"os"
	"testing"

	. "github.com/DnaDiff/forum/src/database"
	_ "github.com/mattn/go-sqlite3"
)

func TestUserDatabase(t *testing.T) {
	// Open a new database connection
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.db")
	defer db.Close()

	// Create a new user
	u := &User{
		// ProfilePicture: "",
		Username:       "janne",
		Age:            5,
		Gender:         "male",
		FirstName:      "Jane",
		LastName:       "Doe",
		Password:       "password456",
		Email:          "janedoe@example.com",
	}
	if err := CreateUser(db, u); err != nil {
		t.Fatal(err)
	}

	// Get the user by username and verify that the fields match
	u2, err := GetUserByUsername(db, "janedoe")
	if err != nil {
		t.Fatal(err)
	}
	if u2.Username != u.Username {
		t.Errorf("expected username %q, got %q", u.Username, u2.Username)
	}
	if u2.Age != u.Age {
		t.Errorf("expected age %d, got %d", u.Age, u2.Age)
	}
	if u2.Gender != u.Gender {
		t.Errorf("expected gender %q, got %q", u.Gender, u2.Gender)
	}
	if u2.FirstName != u.FirstName {
		t.Errorf("expected first name %q, got %q", u.FirstName, u2.FirstName)
	}
	if u2.LastName != u.LastName {
		t.Errorf("expected last name %q, got %q", u.LastName, u2.LastName)
	}
	if u2.Password != u.Password {
		t.Errorf("expected password %q, got %q", u.Password, u2.Password)
	}
	if u2.Email != u.Email {
		t.Errorf("expected email %q, got %q", u.Email, u2.Email)
	}

	// Delete the user and verify that it was deleted
	if err := DeleteUserByUsername(db, "janedoe"); err != nil {
		t.Fatal(err)
	}
	_, err = GetUserByUsername(db, "janedoe")
	if err == nil {
		t.Errorf("expected user not found error, got nil")
	}
	if err != nil && err.Error() != "user: \"janedoe\" not found" {
		t.Errorf("expected user not found error, got %q", err.Error())
	}
}
