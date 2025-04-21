package httpjson

import (
	"context"
	"net/http"

	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/labstack/echo/v4"
)

type CreateAgentRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Location    string `json:"location"`
}

func (h *Handler) CreateAgent(c echo.Context) error {
	ctx := context.Background()

	var req CreateAgentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	agent, err := h.agentSvc.CreateAgent(ctx, app.CreateAgentInput(req))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return handleSuccess(c, http.StatusCreated, newAgentResponse(agent))
}
