package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urusofam/quotesAPI/internal/server/models"
	"github.com/urusofam/quotesAPI/internal/server/repositories"
	"net/http"
	"strconv"
)

type QuoteHandler struct {
	repo repositories.QuoteRepository
}

func NewQuoteHandler(repo repositories.QuoteRepository) *QuoteHandler {
	return &QuoteHandler{repo: repo}
}

func (h *QuoteHandler) PostQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var quote models.Quote
		if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid JSON"))
			return
		}

		if quote.Author == "" || quote.Quote == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Author and Quote are required"))
			return
		}

		h.repo.AddQuote(quote)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Quote added"))
	}
}

func (h *QuoteHandler) DeleteQuoteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
			return
		}

		err = h.repo.DeleteQuoteById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *QuoteHandler) GetRandomQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quote, err := h.repo.GetRandomQuote()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(quote)
	}
}

func (h *QuoteHandler) GetAllQuotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")

		quotes, err := h.repo.GetAllQuotesByAuthor(author)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(quotes)
	}
}
