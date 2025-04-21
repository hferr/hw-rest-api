package httpjson

import (
	"context"
	"net/http"

	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/labstack/echo/v4"
)

type CreateAgentRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Email       string `json:"email" validate:"required,email,max=255"`
	PhoneNumber string `json:"phone_number" validate:"required,max=255"`
	Location    string `json:"location" validate:"required,max=255"`
}

func (h *Handler) CreateAgent(c echo.Context) error {
	ctx := context.Background()

	var req CreateAgentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(req); err != nil {
		return handleValidationError(c, err)
	}

	agent, err := h.agentSvc.CreateAgent(ctx, app.CreateAgentInput(req))
	if err != nil {
		return handleError(c, err)
	}

	return handleSuccess(c, http.StatusCreated, newAgentResponse(agent))
}
