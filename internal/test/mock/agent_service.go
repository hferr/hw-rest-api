package mock

import (
	"context"

	"github.com/hferr/hw-rest-api/internal/app"
)

type AgentService struct {
	CreateAgentFn func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error)
}

func (s *AgentService) CreateAgent(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
	return s.CreateAgentFn(ctx, input)
}
