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

const (
	empty = ""
	tab   = "\t"
)

// PrettyJSON is for pretty printing JSON records
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

// SliceToJSON encodes a slice with JSON records
func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

// SliceFromJSON decodes a serialized slice with JSON records
func SliceFromJSON(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// AddUser is for adding a new user to the database
func AddUser(u User) bool {
	log.Println("Adding user:", u)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(Username, Password, Lastlogin, Admin, Active) values(?,?,?,?,?)")
	if err != nil {
		log.Println("Adduser:", err)
		return false
	}
	stmt.Exec(u.Username, u.Password, u.LastLogin, u.Admin, u.Active)
	return true
}

// UpdateUser allows you to update user name
func UpdateUser(u User) bool {
	log.Println("Updating user:", u)

	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	sqlStatement := `
	UPDATE users
	SET Username = $2, Password = $3, Admin = $4, Active = $5
	WHERE ID = $1;`

	_, err = db.Exec(sqlStatement, u.ID, u.Username, u.Password, u.Admin, u.Active)
	if err != nil {
		log.Println("Update:", err)
		return false
	}

	return true
}

// CreateDatabase initializes the SQLite3 database and adds the admin user
func CreateDatabase() bool {
	log.Println("Writing to SQLite3:", SQLFILE)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	log.Println("Emptying database table.")
	_, _ = db.Exec("DROP TABLE users")

	log.Println("Creating table from scratch.")
	_, err = db.Exec("CREATE TABLE users (ID integer NOT NULL PRIMARY KEY AUTOINCREMENT, Username TEXT, Password TEXT, Lastlogin integer, Admin Bool, Active Bool);")
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("Populating", SQLFILE)
	admin := User{-1, "admin", "admin", time.Now().Unix(), true, false}
	return AddUser(admin)
}

// DeleteUser is for deleting a user defined by ID
func DeleteUser(ID int) bool {
	log.Println("Deleting from SQLite3:", ID)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM users WHERE ID = ?")
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}
	stmt.Exec(ID)
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
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		log.Println(err)
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

	log.Println("All:", all)
	return all
}

// FindUserID is for returning a user record defined by ID
func FindUserID(ID int) User {
	log.Println("Get User Data from SQLite3:", ID)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where ID = $1 \n", ID)
	if err != nil {
		log.Println("Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// FindUserUsername is for returning a user record defined by username
func FindUserUsername(username string) User {
	log.Println("Get User Data from SQLite3:", username)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where Username = $1 \n", username)
	if err != nil {
		log.Println("FindUserUsername Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// ReturnLoggedUsers is for returning all logged in users
func ReturnLoggedUsers() []User {
	log.Println("Reading from SQLite3:", SQLFILE)
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Active = true \n")
	if err != nil {
		log.Println(err)
		return nil
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return []User{}
		}
		temp := User{c1, c2, c3, c4, c5, c6}
		log.Println("temp:", all)
		all = append(all, temp)
	}

	log.Println("Logged in:", all)
	return all
}

// IsUserAdmin determines whether a user is
// an administrator or not
func IsUserAdmin(u UserPass) bool {
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Username = $1 \n", u.Username)
	if err != nil {
		log.Println(err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password && temp.Admin == true {
		return true
	}
	return false
}

func IsUserValid(u UserPass) bool {
	db, err := sql.Open("sqlite3", SQLFILE)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Username = $1 \n", u.Username)
	if err != nil {
		log.Println(err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 bool

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password {
		return true
	}
	return false
}
