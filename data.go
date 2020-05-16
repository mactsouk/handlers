package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var SQLFILE = "/tmp/users.db"

var USERID = 0

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
}

type Input struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type UserPass struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func PrettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

// AddUser is for adding a new user to the database
func AddUser(u User) bool {
	log.Println("Adding user:", u)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(nil)
		return false
	}

	stmt, _ := db.Prepare("INSERT INTO user(ID, Username, Password, Lastlogin, Admin, Active) values(?,?,?,?,?,?)")
	_, _ = stmt.Exec(u.ID, u.Username, u.Password, u.LastLogin, u.Admin, u.Active)

	USERID++
	return true
}

// CreateDatabase initializes the SQLite3 database and adds the admin user
func CreateDatabase() bool {
	log.Println("Writing to SQLite3:", SQLFILE)
	db, err := sql.Open("sqlite3", SQLFILE)

	if err != nil {
		log.Println(nil)
		return false
	}

	log.Println("Emptying database table.")
	_, _ = db.Exec("DROP TABLE users")

	log.Println("Creating table from scratch.")
	_, err = db.Exec("CREATE TABLE users (ID INT, Username STRING, Password STRING, Lastlogin INT64, Admin Bool, Active Bool);")
	if err != nil {
		log.Println(nil)
		return false
	}

	log.Println("Populating", SQLFILE)
	admin := User{USERID, "admin", "admin", time.Now().Unix(), true, true}

	t, _ := PrettyJSON(admin)
	log.Println(t)
	return AddUser(admin)

}

// DeleteUser is for deleting a user defined by ID
func DeleteUser(ID int) bool {

	return true
}

// ReturnAllUsers is for returning all users from database
func ReturnAllUsers() []User {
	log.Println("Reading from SQLite3:", SQLFILE)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return nil
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(nil)
		return nil
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}
	return all
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
