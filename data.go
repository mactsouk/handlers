package handlers

import (
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

func AddUser() {

}

func DeleteUser() {

}

func ReturnAllUsers() {

}

func FindUser() {

}
