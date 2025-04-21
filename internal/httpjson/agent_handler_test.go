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
		inputFn  func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest
		s        *mock.AgentService
	}{
		"success": {
			wantCode: http.StatusCreated,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				return input
			},
			s: &mock.AgentService{
				CreateAgentFn: func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
					return app.Agent{}, nil
				},
			},
		},
		"internal error": {
			wantCode: http.StatusInternalServerError,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				return input
			},
			s: &mock.AgentService{
				CreateAgentFn: func(ctx context.Context, input app.CreateAgentInput) (app.Agent, error) {
					return app.Agent{}, fmt.Errorf("boom")
				},
			},
		},
		"bad request: 'name' is missing": {
			wantCode: http.StatusBadRequest,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				input.Name = ""
				return input
			},
		},
		"bad request: 'email' is missing": {
			wantCode: http.StatusBadRequest,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				input.Email = ""
				return input
			},
		},
		"bad request: 'email' is invalid": {
			wantCode: http.StatusBadRequest,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				input.Email = "invalid_email"
				return input
			},
		},
		"bad request: 'phone_number' is missing": {
			wantCode: http.StatusBadRequest,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				input.PhoneNumber = ""
				return input
			},
		},
		"bad request: 'location' is missing": {
			wantCode: http.StatusBadRequest,
			inputFn: func(input httpjson.CreateAgentRequest) httpjson.CreateAgentRequest {
				input.Location = ""
				return input
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := httpjson.CreateAgentRequest{
				Name:        "test_agent",
				Email:       "test@email.com",
				PhoneNumber: "111-111-1111",
				Location:    "test location",
			}

			input := tc.inputFn(req)

			reqJson, err := json.Marshal(input)
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
