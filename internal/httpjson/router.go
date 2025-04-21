package httpjson

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) NewRouter() http.Handler {
	r := echo.New()

	r.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	return r
}
