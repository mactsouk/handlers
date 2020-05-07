package handlers

import (
	"encoding/json"
	"io"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"user"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"lastlogin"`
	Admin     bool      `json:"admin"`
	Active    bool      `json:"active"`
}

type Input struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// AddUser is for adding a new user to the database
func AddUser(u User) bool {

	return true
}

// DeleteUser is for deleting a user defined by ID
func DeleteUser(ID int) bool {

	return true
}

// ReturnAllUsers is for returning all users from database
func ReturnAllUsers() []User {

	return nil
}

// FindUserID is for returning a user record defined by ID
func FindUserID(ID int) User {

	return User{}
}

// FindUserUsername is for returning a user record defined by username
func FindUserUsername(username string) User {

	return User{}
}

// FromJSON decodes a serialized JSON record
func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON encodes a JSON record
func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
