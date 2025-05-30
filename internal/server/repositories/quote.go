package repositories

import (
	"errors"
	"github.com/urusofam/quotesAPI/internal/server/models"
	"math/rand"
	"time"
)

type QuoteRepository interface {
	AddQuote(quote models.Quote)
	GetAllQuotesByAuthor(author string) ([]models.Quote, error)
	DeleteQuoteById(ID int) error
	GetRandomQuote() (models.Quote, error)
}

type quoteRepository struct {
	quotes []models.Quote
	lastID int
}

func NewQuoteRepository() QuoteRepository {
	return &quoteRepository{quotes: make([]models.Quote, 0), lastID: 0}
}

func (r *quoteRepository) AddQuote(quote models.Quote) {
	quote.ID = r.lastID
	r.lastID++
	r.quotes = append(r.quotes, quote)
}

func (r *quoteRepository) GetAllQuotesByAuthor(author string) ([]models.Quote, error) {
	var quotes []models.Quote

	if len(r.quotes) == 0 {
		return nil, errors.New("no quotes")
	}

	for _, quote := range r.quotes {
		if quote.Author == author || author == "" {
			quotes = append(quotes, quote)
		}
	}
	return quotes, nil
}

func (r *quoteRepository) DeleteQuoteById(ID int) error {
	for i, quote := range r.quotes {
		if quote.ID == ID {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (r *quoteRepository) GetRandomQuote() (models.Quote, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(r.quotes) == 0 {
		return models.Quote{}, errors.New("no quotes")
	}
	idx := rnd.Intn(len(r.quotes))
	return r.quotes[idx], nil
}
