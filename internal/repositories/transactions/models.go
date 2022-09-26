package transactions

import "github.com/google/uuid"

type Transaction struct {
	Id        uuid.UUID `json:"id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	Createdat string    `json:"createdat"`
	Isactive  bool      `json:"isactive"`
}
