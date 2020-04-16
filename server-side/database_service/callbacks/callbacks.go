package callbacks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	datastructures "../data_structures"
	"../utils"
)

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
