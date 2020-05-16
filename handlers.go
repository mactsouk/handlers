package handlers

import (
	"fmt"
	"log"
	"net/http"
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
	if ok {
		log.Println("ID:", id)
	} else {
		log.Println("ID value not set!")
	}

}

// GetAllHandler is for getting all data from the user database
func GetAllHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	Body := "\n"
	for _, d := range ReturnAllUsers() {
		err := d.ToJSON(rw)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(rw, "%s", Body)
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
	if ok {
		log.Println("ID:", id)
	} else {
		log.Println("ID value not set!")
	}

}

// UpdateHandler is for updating the data of an existing user + PUT
func UpdateHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// LoginHandler is for updating the Login time of a user
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// LogoutHandler is for logging out a user
func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

// LoggedUsersHandler returns the list of currently logged in users
func LoggedUsersHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}
