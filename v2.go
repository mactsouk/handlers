package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var IMAGESPATH = "/tmp/files"

type V2Input struct {
	Username  string `json:"username"`
	Upassword string `json:"password"`
	U         User   `json:"load"`
}

func UploadFile(rw http.ResponseWriter, r *http.Request) {
	filename, ok := mux.Vars(r)["filename"]
	if !ok {
		log.Println("filename value not set!")
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(filename)
	saveFile(IMAGESPATH+"/"+filename, rw, r)
}

func saveFile(path string, rw http.ResponseWriter, r *http.Request) {
	log.Println("Saving to", path)
	err := saveToFile(path, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func CreateImageDirectory(d string) error {
	_, err := os.Stat(IMAGESPATH)
	if os.IsNotExist(err) {
		log.Println("Creating:", IMAGESPATH)
		err = os.MkdirAll(IMAGESPATH, 0755)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if err != nil {
		log.Println(err)
		return err
	}

	fileInfo, err := os.Stat(IMAGESPATH)
	mode := fileInfo.Mode()
	if !mode.IsDir() {
		msg := IMAGESPATH + " is not a directory!"
		return errors.New(msg)
	}
	return nil
}

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s from %s", r.RequestURI, r.Host)
		next.ServeHTTP(w, r)
	})
}
