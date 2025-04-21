package httpjson

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FindCustomerAgent(c echo.Context) error {
	ctx := context.Background()

	customerID, err := uuid.Parse(c.Param("customerId"))
	if err != nil {
		return handleValidationError(c, err)
	}

	agentID, err := uuid.Parse(c.Param("agentId"))
	if err != nil {
		return handleValidationError(c, err)
	}

	agent, err := h.customerSvc.FindCustomerAgent(ctx, customerID, agentID)
	if err != nil {
		return handleError(c, err)
	}

	return handleSuccess(c, http.StatusOK, newAgentResponse(agent))
}
