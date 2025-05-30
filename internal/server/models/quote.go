package models

type Quote struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
