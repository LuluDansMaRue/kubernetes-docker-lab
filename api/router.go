package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"./bobba"
	"./database"

	"github.com/gorilla/mux"
)

type messageError struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}

// success struct
type success struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

const (
	idEmptyRemove string = "Id is empty. Can't remove the bubble tea"
	idEmptyDetail string = "Id is empty. Can't get detail of a bubble tea"
)

var (
	db *database.BobbaCon
)

// Default Handler
// Send a small message to the user :)
// Param w http.ResponseWriter
// Param r *http.Request
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello bubble tea server !!"))
}

// Add Bobba
// Add a bubble tea into the database and return the index added
// Param w http.ResponseWriter
// Param r *http.Request
func addBobba(w http.ResponseWriter, r *http.Request) {
	var bobba bobba.Bobba
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bobba)

	// Set JSON header
	w.Header().Set("Content-type", "application/json")

	if err != nil {
		errPayload := buildErrorPayload(err.Error(), 200)
		w.Write(errPayload)
		return
	}

	// call our goroutine
	last, insertErr := runAddRoutine(bobba)

	if insertErr != nil {
		errPayload := buildErrorPayload(insertErr.Error(), 200)
		w.Write(errPayload)
		return
	}

	payload := buildSuccessPayload(last, 200)
	w.Write(payload)
	return
}

// Remove Bobba
// Remove a bobba from the database
// Param w http.ResponseWriter
// Param r *http.Request
func removeBobba(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := getRequestID(vars, "id")

	// Set JSON header
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		errPayload := buildErrorPayload(
			idEmptyRemove,
			200,
		)

		w.Write(errPayload)
		return
	}

}

// Get Bobba Detail
// Get the detail of a bubble tea
// Param http.ResponseWriter w
// Param *http.Request r
func getBobbaDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := getRequestID(vars, "id")

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		errPayload := buildErrorPayload(
			idEmptyDetail,
			200,
		)

		w.Write(errPayload)
		return
	}
}

// Build Error Payload
// Build the error payload message that will be sent by the API
// Param string message
// Param int status
// Return []byte
func buildErrorPayload(message string, status int) []byte {
	payload := messageError{
		Message: message,
		Status:  status,
	}

	json, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return json
}

func buildSuccessPayload(data interface{}, status int) []byte {
	payload := success{
		Data:   data,
		Status: status,
	}

	json, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return json
}

// Get Request ID
// Get the identifier
// Param parmams map[string]string
// Param string identifier
// Return int id
func getRequestID(params map[string]string, identifier string) (int, error) {
	strid := params[identifier]
	id, err := strconv.Atoi(strid)

	if err != nil {
		return 0, errors.New("Unable to convert " + identifier + " to an int")
	}

	return id, nil
}

// InitDB
// Call the initializer of the database
func initDB() {
	instance, err := database.Create()

	if err != nil {
		panic(err)
	}

	db = instance
}

// Creating the routing
func routing() *mux.Router {
	// Initialize the database once
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/bobba/add", addBobba).Methods("POST")
	r.HandleFunc("/bobba/remove/{id}", removeBobba).Methods("DELETE")
	r.HandleFunc("/bobba/{id}", getBobbaDetail).Methods("GET")

	return r
}
