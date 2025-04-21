package app

import (
	"context"

	"github.com/google/uuid"
)

type CustomerService interface {
	FindCustomerAgent(ctx context.Context, customerID, agentID uuid.UUID) (Agent, error)
}

type customerService struct {
	repo Repo
}

func NewCustomerService(r Repo) CustomerService {
	return &customerService{
		repo: r,
	}
}

func (s *customerService) FindCustomerAgent(ctx context.Context, customerID, agentID uuid.UUID) (Agent, error) {
	customer, err := s.repo.FindCustomerByID(ctx, customerID)
	if err != nil {
		if err == ErrCustomerNotFound {
			return Agent{}, err
		}
		return Agent{}, ErrInternal
	}

	agent, err := s.repo.FindAgentByID(ctx, agentID)
	if err != nil {
		if err == ErrAgentNotFound {
			return Agent{}, err
		}
		return Agent{}, ErrInternal
	}

	isCustomerConnectedToAgent, err := s.repo.IsCustomerConnectedToAgent(ctx, customer.ID, agent.ID)
	if err != nil {
		return Agent{}, ErrInternal
	}

	if !isCustomerConnectedToAgent {
		return Agent{}, ErrCustomerAgentConnectionNotFound
	}

	return agent, nil
}
