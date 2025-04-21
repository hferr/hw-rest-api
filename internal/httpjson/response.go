package httpjson

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
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

var errStatusMap = map[error]int{
	app.ErrInternal:                        http.StatusInternalServerError,
	app.ErrAgentNotFound:                   http.StatusNotFound,
	app.ErrCustomerNotFound:                http.StatusNotFound,
	app.ErrCustomerAgentConnectionNotFound: http.StatusNotFound,
}

type errResponse struct {
	Errors []string `json:"errors"`
}

func newErrResponse(errors []string) errResponse {
	return errResponse{Errors: errors}
}

func parseError(err error) []string {
	var errMsgs []string
	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

func handleError(c echo.Context, err error) error {
	statusCode, ok := errStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	return c.JSON(statusCode, newErrResponse(parseError(err)))
}

func handleValidationError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, newErrResponse(parseError(err)))
}

func handleSuccess(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}
