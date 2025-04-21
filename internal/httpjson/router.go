package httpjson

import (
	"net/http"

	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	agentSvc    app.AgentService
	customerSvc app.CustomerService
}

func NewHandler(agentSvc app.AgentService, customerSvc app.CustomerService) *Handler {
	return &Handler{
		agentSvc:    agentSvc,
		customerSvc: customerSvc,
	}
}

func (h *Handler) NewRouter() http.Handler {
	r := echo.New()

	r.Validator = &Validator{validator: NewValidator()}

	r.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	v1 := r.Group("/v1")
	{
		agents := v1.Group("/agents")
		{
			agents.POST("", h.CreateAgent)
			agents.PATCH("/:id", h.UpdateAgent)
		}
		customers := v1.Group("/customers")
		{
			customers.GET("/:customerId/agents/:agentId", h.FindCustomerAgent)
		}
	}

	return r
}
