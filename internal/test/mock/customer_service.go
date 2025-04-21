package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

type CustomerService struct {
	FindCustomerAgentFn func(ctx context.Context, customerID, agentID uuid.UUID) (app.Agent, error)
}

func (s *CustomerService) FindCustomerAgent(ctx context.Context, customerID, agentID uuid.UUID) (app.Agent, error) {
	return s.FindCustomerAgentFn(ctx, customerID, agentID)
}
