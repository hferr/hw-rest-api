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

type CreateAgentInput struct {
	Name        string
	Email       string
	PhoneNumber string
	Location    string
}

type UpdateAgentInput struct {
	Name        *string
	Email       *string
	PhoneNumber *string
	Location    *string
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

func (a Agent) ApplyUpdate(u UpdateAgentInput) Agent {
	if u.Name != nil {
		a.Name = *u.Name
	}

	if u.Email != nil {
		a.Email = *u.Email
	}

	if u.PhoneNumber != nil {
		a.PhoneNumber = *u.PhoneNumber
	}

	if u.Location != nil {
		a.Location = *u.Location
	}

	return a
}
