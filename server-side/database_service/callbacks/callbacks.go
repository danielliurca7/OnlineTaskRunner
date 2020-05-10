package callbacks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	datastructures "../data_structures"
	"../utils"
	"github.com/gorilla/mux"
)

// Authenticate checks if the credentials are valid and, if they are, returns the user data
func Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	var credentials datastructures.Credentials

	if err := json.Unmarshal(body, &credentials); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	if correct, err := utils.VerifyCredentials(credentials); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not access database"))
		return
	} else if !correct {
		log.Println("Incorrect credentials")
		w.WriteHeader(400)
		w.Write([]byte("Incorrect credentials"))
		return
	}

	userdata, err := utils.GetUserData(credentials.Username)

	if b, err := json.Marshal(userdata); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not access database"))
	} else {
		w.Write(b)
	}
}

// GetStudentCourses returns the courses for a student
func GetStudentCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	b, err := utils.GetCourses(name, "student")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get data"))
		return
	}

	w.Write(b)
}

// GetAssistantCourses returns the courses for a assistant
func GetAssistantCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	b, err := utils.GetCourses(name, "assistant")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get data"))
		return
	}

	w.Write(b)
}

// GetProfessorCourses returns the courses for a professor
func GetProfessorCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	b, err := utils.GetCourses(name, "professor")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get data"))
		return
	}

	w.Write(b)
}
