package httpjson_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/hferr/hw-rest-api/internal/httpjson"
	"github.com/hferr/hw-rest-api/internal/test"
	"github.com/hferr/hw-rest-api/internal/test/mock"
)

func TestHandlerCreateAgent(t *testing.T) {
	var testCases = map[string]struct {
		wantCode int
		input    httpjson.CreateAgentRequest
		s        *mock.AgentService
	}{
		"success": {
			wantCode: http.StatusCreated,
			input: httpjson.CreateAgentRequest{
				Name:        "test_agent",
				Email:       "test@email.com",
				PhoneNumber: "111-111-1111",
				Location:    "test location",
			},
			s: &mock.AgentService{
				CreateAgentFn: func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
					return app.Agent{}, nil
				},
			},
		},
		"internal error": {
			wantCode: http.StatusInternalServerError,
			input: httpjson.CreateAgentRequest{
				Name:        "test_agent",
				Email:       "test@email.com",
				PhoneNumber: "111-111-1111",
				Location:    "test location",
			},
			s: &mock.AgentService{
				CreateAgentFn: func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
					return app.Agent{}, fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reqJson, err := json.Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			handler := httpjson.NewHandler(tc.s)
			resp := test.DoHttpRequest(
				handler,
				http.MethodPost,
				"/v1/agents",
				bytes.NewReader(reqJson),
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}
