package app

import "context"

type Repo interface {
	UpsertAgent(ctx context.Context, agent Agent) (Agent, error)
}
