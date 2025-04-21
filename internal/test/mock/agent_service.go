package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

type AgentService struct {
	CreateAgentFn func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error)
	UpdateAgentFn func(ctx context.Context, ID uuid.UUID, input app.UpdateAgentInput) (app.Agent, error)
}

func (s *AgentService) CreateAgent(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
	return s.CreateAgentFn(ctx, input)
}

func (s *AgentService) UpdateAgent(ctx context.Context, ID uuid.UUID, input app.UpdateAgentInput) (app.Agent, error) {
	return s.UpdateAgentFn(ctx, ID, input)
}
