package models

import "github.com/google/uuid"

type Payment struct {
	CorrelationId uuid.UUID `json:"correlationId"`
	Amount        float64   `json:"amount"`
}
