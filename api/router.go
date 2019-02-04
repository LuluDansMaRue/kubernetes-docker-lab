package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type messageError struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}

const (
	idEmptyRemove string = "Id is empty. Can't remove the bubble tea"
	idEmptyDetail string = "Id is empty. Can't get detail of a bubble tea"
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
	var bobba Bobba
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bobba)

	// Set JSON header
	w.Header().Set("Content-type", "application/json")

	if err != nil {
		errPayload := buildErrorPaylaod(err.Error(), 200)
		w.Write(errPayload)
	}
}

// Remove Bobba
// Remove a bobba from the database
// Param w http.ResponseWriter
// Param r *http.Request
func removeBobba(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Set JSON header
	w.Header().Set("Content-type", "application/json")
	if len(id) == 0 {
		errPayload := buildErrorPaylaod(
			idEmptyRemove,
			200,
		)

		w.Write(errPayload)
		return
	}

}

func getBobbaDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-type", "application/json")
	if len(id) == 0 {
		errPayload := buildErrorPaylaod(
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
func buildErrorPaylaod(message string, status int) []byte {
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

func getRequestID(params map[string]string) (int, error) {
	//id := params["id"]

	return 0, nil
}

// Creating the routing
func routing() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/bobba/add", addBobba).Methods("POST")
	r.HandleFunc("/bobba/remove/{id}", removeBobba).Methods("DELETE")
	r.HandleFunc("/bobba/{id}", getBobbaDetail).Methods("GET")

	return r
}
