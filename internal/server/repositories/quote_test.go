package repositories

import (
	"github.com/urusofam/quotesAPI/internal/server/models"
	"testing"
)

func TestAddQuote(t *testing.T) {
	repo := NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "John Doe", Quote: "The Hood"})

	quotes, err := repo.GetAllQuotesByAuthor("John Doe")
	if err != nil {
		t.Error("expected err to be nil, got " + err.Error())
	}
	if len(quotes) != 1 {
		t.Error("Expected 1 quote, got ", len(quotes))
	}
	if quotes[0].Author != "John Doe" {
		t.Error("Expected author to be John Doe, got ", quotes[0].Author)
	}
	if quotes[0].Quote != "The Hood" {
		t.Error("Expected quote to be 'The Hood', got ", quotes[0].Quote)
	}
}

func TestDeleteQuoteById(t *testing.T) {
	repo := NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "John Doe", Quote: "The Hood"})

	err := repo.DeleteQuoteById(0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	quotes, err := repo.GetAllQuotesByAuthor("")
	if err == nil {
		t.Error("expected err to be not nil")
	}
	if len(quotes) != 0 {
		t.Errorf("expected 0 quotes after delete, got %d", len(quotes))
	}
}

func TestDeleteQuoteById_NotExisting(t *testing.T) {
	repo := NewQuoteRepository()
	err := repo.DeleteQuoteById(123)
	if err == nil {
		t.Errorf("expected error on deleting non-existing ID, got nil")
	}
}

func TestGetAllQuotes(t *testing.T) {
	repo := NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "John Doe", Quote: "The Hood"})
	repo.AddQuote(models.Quote{Author: "Big", Quote: "Small"})

	quotes, err := repo.GetAllQuotesByAuthor("John Doe")
	if err != nil {
		t.Error("expected err to be nil, got " + err.Error())
	}
	if len(quotes) != 1 {
		t.Error("Expected 1 quote, got ", len(quotes))
	}

	quotes, err = repo.GetAllQuotesByAuthor("")
	if err != nil {
		t.Error("expected err to be nil, got " + err.Error())
	}
	if len(quotes) != 2 {
		t.Error("Expected 2 quotes, got ", len(quotes))
	}

	quotes, err = repo.GetAllQuotesByAuthor("None")
	if err != nil {
		t.Error("expected err to be nil, got " + err.Error())
	}
	if len(quotes) != 0 {
		t.Error("Expected 0 quotes, got ", len(quotes))
	}
}

func TestGetAllQuotes_Empty(t *testing.T) {
	repo := NewQuoteRepository()
	_, err := repo.GetAllQuotesByAuthor("")
	if err == nil {
		t.Error("expected err to be not nil")
	}
}

func TestGetRandomQuote(t *testing.T) {
	repo := NewQuoteRepository()
	repo.AddQuote(models.Quote{Author: "John Doe", Quote: "The Hood"})
	quote, err := repo.GetRandomQuote()
	if err != nil {
		t.Error("expected err to be nil, got ", err.Error())
	}
	if quote.Author != "John Doe" {
		t.Error("Expected author to be John Doe, got ", quote.Author)
	}
}

func TestGetRandomQuote_Empty(t *testing.T) {
	repo := NewQuoteRepository()
	_, err := repo.GetRandomQuote()
	if err == nil {
		t.Error("Finding an empty quotes should have returned an error")
	}
}
