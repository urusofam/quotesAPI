package repositories

import (
	"errors"
	"github.com/urusofam/quotesAPI/internal/server/models"
)

type QuoteRepository interface {
	AddQuote(quote models.Quote)
	GetAllQuotes() []models.Quote
	GetQuoteByAuthor(author string) (models.Quote, error)
	DeleteQuoteById(ID string) error
}

type quoteRepository struct {
	quotes []models.Quote
	lastID int
}

func NewQuoteRepository() QuoteRepository {
	return &quoteRepository{quotes: make([]models.Quote, 0), lastID: 0}
}

func (r *quoteRepository) AddQuote(quote models.Quote) {
	r.quotes = append(r.quotes, quote)
}

func (r *quoteRepository) GetAllQuotes() []models.Quote {
	return r.quotes
}

func (r *quoteRepository) GetQuoteByAuthor(author string) (models.Quote, error) {
	for _, quote := range r.quotes {
		if quote.Author == author {
			return quote, nil
		}
	}
	return models.Quote{}, errors.New("not found")
}

func (r *quoteRepository) DeleteQuoteById(ID string) error {
	for i, quote := range r.quotes {
		if quote.ID == ID {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
