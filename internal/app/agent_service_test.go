package app_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/hferr/hw-rest-api/internal/test"
	"github.com/hferr/hw-rest-api/internal/test/mock"
)

func TestServiceCreateAgent(t *testing.T) {
	input := app.CreateAgentInput{
		Name:        "test_agent",
		Email:       "test@email.com",
		PhoneNumber: "111-111-1111",
		Location:    "test location",
	}

	wantAgent := app.NewAgent(
		input.Name,
		input.Email,
		input.PhoneNumber,
		input.Location,
	)

	var testCases = map[string]struct {
		wantErr error
		input   app.CreateAgentInput
		repo    *mock.Repo
	}{
		"success": {
			wantErr: nil,
			input:   input,
			repo: &mock.Repo{
				UpsertAgentFn: func(ctx context.Context, agent app.Agent) (app.Agent, error) {
					return wantAgent, nil
				},
			},
		},
		"internal error": {
			wantErr: app.ErrInternal,
			input:   input,
			repo: &mock.Repo{
				UpsertAgentFn: func(ctx context.Context, agent app.Agent) (app.Agent, error) {
					return app.Agent{}, app.ErrInternal
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			agentSvc := app.NewAgentService(tc.repo)
			got, err := agentSvc.CreateAgent(context.Background(), tc.input)
			if err != nil {
				if err != tc.wantErr {
					t.Errorf("expected error to be: %v, got: %v", tc.wantErr, err)
				}
			}

			if err == nil {
				if got.ID == uuid.Nil {
					t.Errorf("expected agent ID to be present")
				}

				if got.Name != wantAgent.Name {
					t.Errorf("%s: expected %s but got %s", "Name", wantAgent.Name, got.Name)
				}

				if got.Email != wantAgent.Email {
					t.Errorf("%s: expected %s but got %s", "Email", wantAgent.Email, got.Email)
				}

				if got.PhoneNumber != wantAgent.PhoneNumber {
					t.Errorf("%s: expected %s but got %s", "PhoneNumber", wantAgent.PhoneNumber, got.PhoneNumber)
				}

				if got.Location != wantAgent.Location {
					t.Errorf("%s: expected %s but got %s", "Email", wantAgent.Location, got.Location)
				}
			}
		})
	}
}

func TestServiceUpdateAgent(t *testing.T) {
	originalAgent := app.NewAgent(
		"test_agent",
		"test@email.com",
		"111-111-1111",
		"test location",
	)

	var testCases = map[string]struct {
		wantErr error
		ID      uuid.UUID
		input   app.UpdateAgentInput
		repo    *mock.Repo
	}{
		"success": {
			ID:      originalAgent.ID,
			input:   app.UpdateAgentInput{Name: test.Ptr("updated_agent")},
			wantErr: nil,
			repo: &mock.Repo{
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return originalAgent, nil
				},
				UpsertAgentFn: func(ctx context.Context, agent app.Agent) (app.Agent, error) {
					return agent, nil
				},
			},
		},
		"agent not found": {
			wantErr: app.ErrAgentNotFound,
			ID:      uuid.New(),
			input:   app.UpdateAgentInput{Name: test.Ptr("updated_agent")},
			repo: &mock.Repo{
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return app.Agent{}, app.ErrAgentNotFound
				},
			},
		},
		"internal error on find agent": {
			wantErr: app.ErrInternal,
			ID:      originalAgent.ID,
			input:   app.UpdateAgentInput{Name: test.Ptr("updated_agent")},
			repo: &mock.Repo{
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return app.Agent{}, app.ErrInternal
				},
			},
		},
		"internal error on upsert agent": {
			wantErr: app.ErrInternal,
			ID:      originalAgent.ID,
			input:   app.UpdateAgentInput{Name: test.Ptr("updated_agent")},
			repo: &mock.Repo{
				FindAgentByIDFn: func(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
					return originalAgent, nil
				},
				UpsertAgentFn: func(ctx context.Context, agent app.Agent) (app.Agent, error) {
					return app.Agent{}, app.ErrInternal
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			agentSvc := app.NewAgentService(tc.repo)
			got, err := agentSvc.UpdateAgent(context.Background(), tc.ID, tc.input)
			if err != nil {
				if err != tc.wantErr {
					t.Fatalf("expected error to be: %v, got: %v", tc.wantErr, err)
				}
			}

			if err == nil {
				if got.ID != tc.ID {
					t.Fatalf("wanted agent ID: %s, got: %s", tc.ID.String(), got.ID.String())
				}
				if got.Name != *tc.input.Name {
					t.Fatalf("wanted agent name: %s, got: %s", *tc.input.Name, got.Name)
				}
			}
		})
	}
}
