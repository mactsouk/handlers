package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type V2Input struct {
	Username  string `json:"username"`
	Upassword string `json:"password"`
	U         User   `json:"load"`
}

var IMAGESPATH string

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

func saveToFile(path string, contents io.Reader) error {
	_, err := os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			log.Println("Error deleting", path)
			return err
		}
	} else if !os.IsNotExist(err) {
		log.Println("Unexpected error:", err)
		return err
	}

	// If everything is OK, create the file
	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	n, err := io.Copy(f, contents)
	if err != nil {
		return err
	}
	log.Println("Bytes written:", n)
	return nil
}

func CreateImageDirectory() error {
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
