package httpjson

import (
	"net/http"

	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	agentSvc app.AgentService
}

func NewHandler(agentSvc app.AgentService) *Handler {
	return &Handler{
		agentSvc: agentSvc,
	}
}

func (h *Handler) NewRouter() http.Handler {
	r := echo.New()

	r.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	v1 := r.Group("/v1")
	{
		agents := v1.Group("/agents")
		{
			agents.POST("", h.CreateAgent)
		}
	}

	return r
}
