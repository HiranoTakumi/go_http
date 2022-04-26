package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	h := NewHandler()
	t.Run("sample_test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjMxMjgxNzcsImlhdCI6MTY0NzU3NjE3NywidG9rZW5UeXBlIjoicmVmcmVzaCIsInVzZXJOYW1lIjoidGVzdCJ9.gGaXmyNFMW-rB_5L7huNodlslc7uOsX8ylOLUwrgqbA")

		res := httptest.NewRecorder()
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
		e.GET("/", h.Sample)
		e.Use(middleware.JWT([]byte("secret")))
		e.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, `{"admin":true,"name":"John Doe","sub":"1234567890"}`+"\n", res.Body.String())
	})
}
