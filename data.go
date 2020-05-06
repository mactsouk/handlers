package handlers

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin time   `json:"lastlogin"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
}
