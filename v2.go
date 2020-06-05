package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type V2Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
	U        User   `json:"load"`
}

var IMAGESPATH string

func AddHandlerV2(rw http.ResponseWriter, r *http.Request) {
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

	var load = V2Input{}
	err = json.Unmarshal(d, &load)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(load)

	u := UserPass{load.Username, load.Password}
	if !IsUserAdmin(u) {
		log.Println("Command issued by non-admin user:", u.Username)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser := load.U
	result := AddUser(newUser)
	if !result {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func LoginHandlerV2(rw http.ResponseWriter, r *http.Request) {
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

	var load = V2Input{}
	err = json.Unmarshal(d, &load)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func LogoutHandlerV2(rw http.ResponseWriter, r *http.Request) {
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

	var load = V2Input{}
	err = json.Unmarshal(d, &load)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetAllHandlerV2(rw http.ResponseWriter, r *http.Request) {
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

	var load = V2Input{}
	err = json.Unmarshal(d, &load)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = UserPass{load.Username, load.Password}
	if !IsUserAdmin(user) {
		log.Println("User", user, "does not exist!")
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

// GetAllHandlerUpdated is for `/v1/getall`.
// The older version had a bug as it was using `IsUserValid` instead of `IsUserAdmin`.
func GetAllHandlerUpdated(rw http.ResponseWriter, r *http.Request) {
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

	if !IsUserAdmin(user) {
		log.Println("User", user, "is not Admin!")
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	err = SliceToJSON(ReturnAllUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
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

func CreateImageDirectory(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		log.Println("Creating:", d)
		err = os.MkdirAll(d, 0755)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if err != nil {
		log.Println(err)
		return err
	}

	fileInfo, err := os.Stat(d)
	mode := fileInfo.Mode()
	if !mode.IsDir() {
		msg := d + " is not a directory!"
		return errors.New(msg)
	}
	return nil
}

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s from %s using %s method", r.RequestURI, r.Host, r.Method)
		next.ServeHTTP(w, r)
	})
}

// Generating Random Strings with a Given length
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomPassword generates random strings of given length
func RandomPassword(l int) string {
	Password := ""
	rand.Seed(time.Now().Unix())
	MIN := 0
	MAX := 94
	startChar := "!"
	i := 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		Password += newChar
		if i == l {
			break
		}
		i++
	}
	return Password
}
