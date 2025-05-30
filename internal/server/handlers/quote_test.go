package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urusofam/quotesAPI/internal/server/models"
	"github.com/urusofam/quotesAPI/internal/server/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostQuote_Success(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).PostQuote()

	body := bytes.NewBufferString(`{"author":"A","quote":"Q"}`)
	req := httptest.NewRequest(http.MethodPost, "/quotes", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
}

func TestPostQuote_InvalidJSON(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).PostQuote()

	body := bytes.NewBufferString(`{"author":123}`)
	req := httptest.NewRequest(http.MethodPost, "/quotes", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestPostQuote_EmptyFields(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).PostQuote()

	body := bytes.NewBufferString(`{"author":"","quote":""}`)
	req := httptest.NewRequest(http.MethodPost, "/quotes", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGetAllQuotes(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "A", Quote: "Q"})
	repo.AddQuote(models.Quote{Author: "B", Quote: "Q2"})
	handler := NewQuoteHandler(repo).GetAllQuotes()

	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var quotes []models.Quote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if len(quotes) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(quotes))
	}
}

func TestGetAllQuotes_NoQuotes(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).GetAllQuotes()
	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGetRandomQuote(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "A", Quote: "Q"})
	handler := NewQuoteHandler(repo).GetRandomQuote()
	req := httptest.NewRequest(http.MethodGet, "/quotes/random", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var q models.Quote
	if err := json.NewDecoder(resp.Body).Decode(&q); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if q.Author != "A" || q.Quote != "Q" {
		t.Errorf("unexpected quote: %+v", q)
	}
}

func TestGetRandomQuote_Empty(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).GetRandomQuote()
	req := httptest.NewRequest(http.MethodGet, "/quotes/random", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestDeleteQuoteById_Success(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "A", Quote: "Q"})
	handler := NewQuoteHandler(repo).DeleteQuoteById()
	req := httptest.NewRequest(http.MethodDelete, "/quotes/0", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "0"})
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestDeleteQuoteById_NotFound(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).DeleteQuoteById()
	req := httptest.NewRequest(http.MethodDelete, "/quotes/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestDeleteQuoteById_InvalidID(t *testing.T) {
	repo := repositories.NewQuoteRepository()
	handler := NewQuoteHandler(repo).DeleteQuoteById()
	req := httptest.NewRequest(http.MethodDelete, "/quotes/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}
