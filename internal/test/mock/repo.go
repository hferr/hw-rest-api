package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

type Repo struct {
	FindAgentByIDFn func(ctx context.Context, ID uuid.UUID) (app.Agent, error)
	UpsertAgentFn   func(ctx context.Context, agent app.Agent) (app.Agent, error)
}

func (r *Repo) FindAgentByID(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
	return r.FindAgentByIDFn(ctx, ID)
}

func (r *Repo) UpsertAgent(ctx context.Context, agent app.Agent) (app.Agent, error) {
	return r.UpsertAgentFn(ctx, agent)
}
