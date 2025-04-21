package app

import (
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID          uuid.UUID
	Name        string
	Email       string
	PhoneNumber string
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAgent(name, email, phoneNumber, location string) Agent {
	return Agent{
		ID:          uuid.New(),
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Location:    location,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type CreateAgentInput struct {
	Name        string
	Email       string
	PhoneNumber string
	Location    string
}
