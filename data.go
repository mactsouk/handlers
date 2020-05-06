package handlers

import "time"

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
