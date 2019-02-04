package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func addBobba(w http.ResponseWriter, r *http.Request) {

}

func removeBobba(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

func getBobbaDetail(w http.ResponseWriter, r *http.Request) {

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
