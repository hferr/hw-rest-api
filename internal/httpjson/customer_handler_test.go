package httpjson_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/hferr/hw-rest-api/internal/httpjson"
	"github.com/hferr/hw-rest-api/internal/test"
	"github.com/hferr/hw-rest-api/internal/test/mock"
)

func TestHandlerFindCustomerAgent(t *testing.T) {
	var testCases = map[string]struct {
		wantCode   int
		customerID string
		agentID    string
		s          *mock.CustomerService
	}{
		"success": {
			wantCode:   http.StatusOK,
			customerID: uuid.New().String(),
			agentID:    uuid.New().String(),
			s: &mock.CustomerService{
				FindCustomerAgentFn: func(ctx context.Context, customerID, agentID uuid.UUID) (app.Agent, error) {
					return app.Agent{}, nil
				},
			},
		},
		"internal error": {
			wantCode:   http.StatusInternalServerError,
			customerID: uuid.New().String(),
			agentID:    uuid.New().String(),
			s: &mock.CustomerService{
				FindCustomerAgentFn: func(ctx context.Context, customerID, agentID uuid.UUID) (app.Agent, error) {
					return app.Agent{}, fmt.Errorf("boom")
				},
			},
		},
		"bad request: invalid 'customerID' in query param": {
			wantCode:   http.StatusBadRequest,
			customerID: "invalid",
			agentID:    uuid.New().String(),
		},
		"bad request: invalid 'agentID' in query param": {
			wantCode:   http.StatusBadRequest,
			customerID: uuid.New().String(),
			agentID:    "invalid",
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			handler := httpjson.NewHandler(nil, tc.s)
			resp := test.DoHttpRequest(
				handler,
				http.MethodGet,
				fmt.Sprintf("/v1/customers/%s/agents/%s", tc.customerID, tc.agentID),
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Errorf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}
