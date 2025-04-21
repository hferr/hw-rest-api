package httpjson

import (
	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/labstack/echo/v4"
)

type agentResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Location    string    `json:"location"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func newAgentResponse(agent app.Agent) agentResponse {
	return agentResponse{
		ID:          agent.ID,
		Name:        agent.Name,
		Email:       agent.Email,
		PhoneNumber: agent.PhoneNumber,
		Location:    agent.Location,
		CreatedAt:   agent.CreatedAt.String(),
		UpdatedAt:   agent.UpdatedAt.String(),
	}
}

func handleSuccess(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}
