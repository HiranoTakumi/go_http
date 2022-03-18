package handlers

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/golang-jwt/jwt/v4"
  // jwt/v4をインポートするとPanicする。echo内ではv4としていしていないから。
)

func Sample(c echo.Context) error {
  token := c.Get("user").(*jwt.Token)
  return c.JSON(http.StatusOK, token.Claims)
}