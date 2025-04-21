package app

import (
	"context"

	"github.com/google/uuid"
)

type AgentService interface {
	CreateAgent(ctx context.Context, input CreateAgentInput) (Agent, error)
	UpdateAgent(ctx context.Context, ID uuid.UUID, input UpdateAgentInput) (Agent, error)
}

type agentService struct {
	repo Repo
}

func NewAgentService(r Repo) AgentService {
	return &agentService{
		repo: r,
	}
}

func (s *agentService) CreateAgent(ctx context.Context, input CreateAgentInput) (Agent, error) {
	agent := NewAgent(
		input.Name,
		input.Email,
		input.PhoneNumber,
		input.Location,
	)

	agent, err := s.repo.UpsertAgent(ctx, agent)
	if err != nil {
		return Agent{}, err
	}

	return agent, nil
}

func (s *agentService) UpdateAgent(ctx context.Context, ID uuid.UUID, input UpdateAgentInput) (Agent, error) {
	original, err := s.repo.FindAgentByID(ctx, ID)
	if err != nil {
		return Agent{}, err
	}

	updated := original.ApplyUpdate(input)

	agent, err := s.repo.UpsertAgent(ctx, updated)
	if err != nil {
		return Agent{}, err
	}

	return agent, nil
}
