package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRefresh(t *testing.T) {
	t.Run("auth_test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/restricted/refresh", nil)
		req.Header.Set(echo.HeaderAuthorization, "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjMxMjgxNzcsImlhdCI6MTY0NzU3NjE3NywidG9rZW5UeXBlIjoicmVmcmVzaCIsInVzZXJOYW1lIjoidGVzdCJ9.gGaXmyNFMW-rB_5L7huNodlslc7uOsX8ylOLUwrgqbA")

		res := httptest.NewRecorder()
		e := echo.New()
		e.GET("/restricted/refresh", Refresh)
		e.Use(middleware.JWT([]byte("secret")))
		e.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})
}
