package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/xid"

	"../data_structures/credentials"
	"../data_structures/request"
	"../utils"
)

var db *sql.DB
var tokens = make(map[string][]byte)

func authenticate(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var credentials credentials.Credentials
	json.Unmarshal(body, &credentials)

	log.Println("Authenticate request for user", credentials.Username)

	_, err := db.Exec("set @valid = Login(?, ?);", credentials.Username, credentials.PasswordHash)
	utils.CheckError(err)

	row := db.QueryRow("select @valid")

	valid := false
	row.Scan(&valid)

	if !valid {
		return
	}

	h := sha1.New()

	h.Write([]byte(credentials.Username))

	token := append([]byte(hex.EncodeToString(h.Sum(nil))), []byte(xid.New().String())...)

	tokens[credentials.Username] = token

	utils.WriteResponse(w, 200, token)
}

func validate(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var request request.Request
	json.Unmarshal(body, &request)

	tokenString := string(request.Request)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return tokens[request.Username], nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["request"])

		// TODO Check permisions
	} else {
		log.Println(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var request request.Request
	json.Unmarshal(body, &request)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/university")
	utils.CheckError(err)
	defer db.Close()

	if string(body) == "insert" {
		_, err := db.Exec("insert into Users(UserName, FirstName, LastName) values ('daniel.liurca', 'Daniel', 'Liurca');")
		utils.CheckError(err)
	} else if string(body) == "select" {
		var username string
		var firstname string
		var lastname string

		rows, err := db.Query("select UserName, FirstName, LastName from Users;")
		if err != nil {
			fmt.Println(err)
		}

		rows.Next()
		err = rows.Scan(&username, &firstname, &lastname)
		utils.CheckError(err)

		fmt.Println(username, firstname, lastname)
	} else if string(body) == "call" {
		userName := "daniel.liurca"
		courseName := "CPL"
		schoolyear := 2019
		assignmentName := "Tema 1"

		_, err := db.Exec("set @valid = IsValid(?, ?, ?, ?);", userName, courseName, schoolyear, assignmentName)
		utils.CheckError(err)

		row := db.QueryRow("select @valid")

		var valid bool
		switch err := row.Scan(&valid); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(valid)
		default:
			utils.CheckError(err)
		}
	}
}

func main() {
	r := mux.NewRouter()

	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/university")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/api/test", test).Methods("POST")

	r.HandleFunc("/api/authenticate", authenticate).Methods("POST")
	r.HandleFunc("/api/validate", validate).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
