package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

type Repo struct {
	FindAgentByIDFn func(ctx context.Context, ID uuid.UUID) (app.Agent, error)
	UpsertAgentFn   func(ctx context.Context, agent app.Agent) (app.Agent, error)

	FindCustomerByIDFn           func(ctx context.Context, ID uuid.UUID) (app.Customer, error)
	IsCustomerConnectedToAgentFn func(ctx context.Context, customerID, agentID uuid.UUID) (bool, error)
}

func (r *Repo) FindAgentByID(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
	return r.FindAgentByIDFn(ctx, ID)
}

func (r *Repo) UpsertAgent(ctx context.Context, agent app.Agent) (app.Agent, error) {
	return r.UpsertAgentFn(ctx, agent)
}

func (r *Repo) FindCustomerByID(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
	return r.FindCustomerByIDFn(ctx, ID)
}

func (r *Repo) IsCustomerConnectedToAgent(ctx context.Context, customerID, agentID uuid.UUID) (bool, error) {
	return r.IsCustomerConnectedToAgentFn(ctx, customerID, agentID)
}
