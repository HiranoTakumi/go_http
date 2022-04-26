package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthcheck handles GET /helthcheck
func (h *Handler) Healthcheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
