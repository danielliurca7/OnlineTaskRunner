package compute

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	datastructures "../../data_structures"
	"../../utils"
)

// ResponseTime is the metric that we are monitoring
var ResponseTime = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "response_time",
		Name:      "compute_service",
		Help:      "Histogram for the response time for the compute service in miliseconds",
	})

// BuildImage verifies the the validity of the token and the request body
// If token is valid, redirects the request to compute service
func BuildImage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var buildBody datastructures.BuildBody
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &buildBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(buildBody.Workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://compute:5000/api/build", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)

	elapsed := float64(time.Now().Sub(start)) / float64(time.Millisecond)
	ResponseTime.Observe(elapsed)
}

// RunContainer verifies the the validity of the token and the request body
// If token is valid, redirects the request to compute service
func RunContainer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var workspace datastructures.Workspace
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &workspace); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://compute:5000/api/run", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)

	elapsed := float64(time.Now().Sub(start)) / float64(time.Millisecond)
	ResponseTime.Observe(elapsed)
}

// StopContainer verifies the the validity of the token and the request body
// If token is valid, redirects the request to compute service
func StopContainer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var workspace datastructures.Workspace
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &workspace); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://compute:5000/api/stop", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)

	elapsed := float64(time.Now().Sub(start)) / float64(time.Millisecond)
	ResponseTime.Observe(elapsed)
}

// ExecContainer verifies the the validity of the token and the request body
// If token is valid, redirects the request to compute service
func ExecContainer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var execBody datastructures.ExecBody
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &execBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(execBody.Workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://compute:5000/api/exec", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)

	elapsed := float64(time.Now().Sub(start)) / float64(time.Millisecond)
	ResponseTime.Observe(elapsed)
}
