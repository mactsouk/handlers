package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DefaultHandler is for handling /
func DefaultHandler(rw http.ResponseWriter, r *http.Request) {

}

// TimeHandler is for handling /time
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func AddHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func GetAllHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func GetHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func UpdateHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	for k, v := range mux.Vars(r) {
		log.Println("K:", k, "v:", v)
	}

}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}
