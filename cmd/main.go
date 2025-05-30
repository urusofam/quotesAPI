package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	router.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/quotes/random", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/quotes/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
