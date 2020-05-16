package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// DefaultHandler is for handling /
func DefaultHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := r.URL.Path + " is not supported. Thanks for visiting!\n"
	fmt.Fprintf(rw, "%s", Body)
}

// MethodNotAllowedHandler is executed when the method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := "Method not allowed!\n"
	fmt.Fprintf(rw, "%s", Body)
}

// TimeHandler is for handling /time
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	rw.WriteHeader(http.StatusOK)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	fmt.Fprintf(rw, "%s", Body)
}

// AddHandler is for adding a new user
func AddHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// DeleteHandler is for deleting an existing user + DELETE
func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = UserPass{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAdmin(user) {
		log.Println("User", user.Username, "is not admin!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id", err)
		return
	}

	t := FindUserID(intID)
	if t.Username != "" {
		log.Println("About to delete:", t)
		deleted := DeleteUser(intID)
		if deleted {
			log.Println("User deleted:", id)
			rw.WriteHeader(http.StatusOK)
			return
		} else {
			log.Println("Cannot delete user:", id)
			rw.WriteHeader(http.StatusNotFound)
		}
	}

	rw.WriteHeader(http.StatusNotFound)
}

// GetAllHandler is for getting all data from the user database
func GetAllHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = UserPass{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SliceToJSON(ReturnAllUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

// GetIDHandler returns the ID of an existing user
func GetIDHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// GetUserDataHandler + GET returns the full record of a user
func GetUserDataHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := FindUserID(intID)
	if t.Username != "" {
		err := t.ToJSON(rw)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
	} else {
		log.Println("User not found:", id)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

// UpdateHandler is for updating the data of an existing user + PUT
func UpdateHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// LoginHandler is for updating the LastLogin time of a user
// And changing the Active field to true
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = UserPass{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Input user:", user)

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := FindUserUsername(user.Username)
	log.Println("Logging in:", t)

	t.LastLogin = time.Now().Unix()
	t.Active = 1
	if UpdateUser(t) {
		log.Println("User updated:", t)
	} else {
		log.Println("Update failed:", t)
		rw.WriteHeader(http.StatusBadRequest)
	}
}

// LogoutHandler is for logging out a user
// And changing the Active field to false
func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = UserPass{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := FindUserUsername(user.Username)
	log.Println("Logging out:", t)
	t.Active = 0
	if UpdateUser(t) {
		log.Println("User updated:", t)
	} else {
		log.Println("Update failed:", t)
		rw.WriteHeader(http.StatusBadRequest)
	}

}

// LoggedUsersHandler returns the list of currently logged in users
func LoggedUsersHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = UserPass{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SliceToJSON(ReturnLoggedUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
