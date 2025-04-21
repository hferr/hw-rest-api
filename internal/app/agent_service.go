package app

import "context"

type AgentService interface {
	CreateAgent(ctx context.Context, input CreateAgentInput) (Agent, error)
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
