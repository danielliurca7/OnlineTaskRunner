package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	datastructures "../../data_structures"
	"../../utils"
	"github.com/gorilla/mux"
)

// Authenticate receives a username and a password hash and return a jwt
func Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	data, status, err := utils.MakeRequest("GET", "http://database:6000/api/authenticate", body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(status)
		w.Write([]byte("Could not verify credentials"))
		return
	}

	var userdata datastructures.UserData
	if err := json.Unmarshal(data, &userdata); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not verify credentials"))
		return
	}

	token, err := utils.BuildToken(userdata)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not build token"))
		return
	}

	w.Write(token)
}

// GetStudentCourses returns the courses for a student
func GetStudentCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	token := r.Header.Get("Authorization")

	userdata, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata.Name != name {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("GET", "http://database:6000/api/student/"+name, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(status)
		w.Write([]byte("Could not verify credentials"))
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}

// GetAssistantCourses returns the courses for a student
func GetAssistantCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	token := r.Header.Get("Authorization")

	userdata, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata.Name != name {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("GET", "http://database:6000/api/assistant/"+name, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(status)
		w.Write([]byte("Could not verify credentials"))
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}

// GetProfessorCourses returns the courses for a student
func GetProfessorCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	token := r.Header.Get("Authorization")

	userdata, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata.Name != name {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("GET", "http://database:6000/api/professor/"+name, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(status)
		w.Write([]byte("Could not verify credentials"))
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}
