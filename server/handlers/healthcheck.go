package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Healthcheck handles GET /helthcheck
func (h *Handler) Healthcheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
