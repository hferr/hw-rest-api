package app

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	FindAgentByID(ctx context.Context, ID uuid.UUID) (Agent, error)
	UpsertAgent(ctx context.Context, agent Agent) (Agent, error)

	FindCustomerByID(ctx context.Context, ID uuid.UUID) (Customer, error)
	IsCustomerConnectedToAgent(ctx context.Context, customerID, agentID uuid.UUID) (bool, error)
}
