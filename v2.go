package handlers

import "net/http"

type V2Input struct {
	Username  string `json:"username"`
	Upassword string `json:"password"`
	U         User   `json:"load"`
}

func uploadFile(rw http.ResponseWriter, r *http.Request) {

}

func sendFile(rw http.ResponseWriter, r *http.Request) {

}
