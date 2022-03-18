package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	userIDKey = "USER_ID_KEY"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter missing."})
	}

	if username == "test" && password == "test" {
		at, rt, err := generateToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  at,
			"refreshToken": rt,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "authentication failed."})
}

func Refresh(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenType := claims["tokenType"].(string)
	name := claims["userName"].(string)
	if tokenType == "refresh" && name == "test" {
		at, rt, err := generateToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  at,
			"refreshToken": rt,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": middleware.ErrJWTMissing.Message.(string)})
}

func generateToken() (string, string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	// set access token claims
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["userName"] = "test"
	accessClaims["id"] = "test_id"
	accessClaims["tokenType"] = "access"
	accessClaims["iat"] = time.Now().Unix()
	accessClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// set refresh token claims
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["userName"] = "test"
	accessClaims["id"] = "test_id"
	refreshClaims["tokenType"] = "refresh"
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 180).Unix()

	// generate encoded token and send it as response
	at, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}
	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}
	return at, rt, nil
}
