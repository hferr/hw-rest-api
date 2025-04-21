package mock

import (
	"context"

	"github.com/hferr/hw-rest-api/internal/app"
)

type Repo struct {
	UpsertAgentFn func(ctx context.Context, agent app.Agent) (app.Agent, error)
}

func (r *Repo) UpsertAgent(ctx context.Context, agent app.Agent) (app.Agent, error) {
	return r.UpsertAgentFn(ctx, agent)
}
