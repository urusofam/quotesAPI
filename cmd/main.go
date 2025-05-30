package main

import (
	"github.com/gorilla/mux"
	"github.com/urusofam/quotesAPI/internal/server/handlers"
	"github.com/urusofam/quotesAPI/internal/server/repositories"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	repoHandler := repositories.NewQuoteRepository()
	handler := handlers.NewQuoteHandler(repoHandler)

	router.HandleFunc("/quotes", handler.PostQuote()).Methods("POST")
	router.HandleFunc("/quotes", handler.GetAllQuotes()).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuote()).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuoteById()).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
