package app

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID             uuid.UUID
	Name           string
	Email          string
	PhoneNumber    string
	CurrentAddress string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
