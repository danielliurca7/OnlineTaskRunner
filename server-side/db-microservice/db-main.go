package main

import (
	"encoding/json"
	"log"
	"net/http"

	"../data_structures/token"
	"../utils"
	"github.com/gorilla/mux"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Authenticate request for user", string(body))

	// Make request to database or cs.curs api

	err := error(nil)
	token := token.Token{
		Token:        "db187a1d6e4d1419b2e1110939f809e3",
		PrivateToken: "BVX0hJVCUB3vJ0FnREyGx6cHt3NwWSM43IrSiLNbcCSY29iGQxjSxQObrFU7YcG2",
	}

	b, err := json.Marshal(token)
	utils.CheckError(err)

	utils.WriteResponse(w, 200, b)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/authenticate", authenticate).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
