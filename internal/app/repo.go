package app

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	FindAgentByID(ctx context.Context, ID uuid.UUID) (Agent, error)
	UpsertAgent(ctx context.Context, agent Agent) (Agent, error)
}
