package app_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/hferr/hw-rest-api/internal/test/mock"
)

func TestServiceFindCustomerAgent(t *testing.T) {
	wantAgent := app.NewAgent(
		"test_agent",
		"test_email",
		"111-111-1111",
		"test location",
	)

	var testCases = map[string]struct {
		wantErr    error
		wantAgent  app.Agent
		customerID uuid.UUID
		agentID    uuid.UUID
		repo       *mock.Repo
	}{
		"success": {
			wantErr:    nil,
			wantAgent:  wantAgent,
			customerID: uuid.New(),
			agentID:    uuid.New(),
			repo: &mock.Repo{
				FindCustomerByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
					return app.Customer{}, nil
				},
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return wantAgent, nil
				},
				IsCustomerConnectedToAgentFn: func(ctx context.Context, customerID, agentID uuid.UUID) (bool, error) {
					return true, nil
				},
			},
		},
		"customer not found": {
			wantErr:    app.ErrCustomerNotFound,
			wantAgent:  wantAgent,
			customerID: uuid.New(),
			agentID:    uuid.New(),
			repo: &mock.Repo{
				FindCustomerByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
					return app.Customer{}, app.ErrCustomerNotFound
				},
			},
		},
		"agent not found": {
			wantErr:    app.ErrAgentNotFound,
			wantAgent:  wantAgent,
			customerID: uuid.New(),
			agentID:    uuid.New(),
			repo: &mock.Repo{
				FindCustomerByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
					return app.Customer{}, nil
				},
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return app.Agent{}, app.ErrAgentNotFound
				},
			},
		},
		"agent is not connected to customer": {
			wantErr:    app.ErrCustomerAgentConnectionNotFound,
			wantAgent:  wantAgent,
			customerID: uuid.New(),
			agentID:    uuid.New(),
			repo: &mock.Repo{
				FindCustomerByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
					return app.Customer{}, nil
				},
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return wantAgent, nil
				},
				IsCustomerConnectedToAgentFn: func(ctx context.Context, customerID, agentID uuid.UUID) (bool, error) {
					return false, nil
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service := app.NewCustomerService(tc.repo)
			got, err := service.FindCustomerAgent(context.Background(), tc.customerID, tc.agentID)

			if err != tc.wantErr {
				t.Errorf("expected error %v, got: %v", tc.wantErr, err)
			}

			if tc.wantErr == nil {
				if got.ID != wantAgent.ID {
					t.Errorf("%s: expected %s, got %s", "ID", wantAgent.ID.String(), got.ID.String())
				}
			}
		})
	}
}
