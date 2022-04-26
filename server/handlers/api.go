package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	// "strings"

	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func (h *Handler) LateResponse(c echo.Context) error {
	second := c.QueryParam("second")
	sleepTime, err := strconv.Atoi(second)
	if err != nil {
		sleepTime = 0
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	log.Println(c.RealIP())
	log.Println(c.Request().Header.Get(echo.HeaderXRealIP))
	log.Println(c.Request().Header.Get(echo.HeaderXForwardedFor))
	log.Println(echo.ExtractIPDirect()(c.Request()))
	log.Println(c.Request().Header.Get(echo.HeaderAuthorization))
	resp := map[string]int{"sleepTime": sleepTime}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListApi(c echo.Context) error {
	res := Persons{}
	res.Persons = []Person{}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) CookieApi(c echo.Context) error {
	secure := true
	httpOnly := false
	accToken := &http.Cookie{
		Name:     "xBaasAccessToken",
		Value:    "token",
		Expires:  time.Now().Add(time.Hour * 3),
		Secure:   secure,
		HttpOnly: httpOnly,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(accToken)
	log.Println(c.Response().Header().Get("Set-Cookie"))
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) ValidCookie(c echo.Context) error {
	// cookie := c.Request().Cookies()
	// // if err != nil {
	// // 	return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	// // }
	// log.Println(cookie.Name)
	// log.Println(cookie.Value)
	// log.Println(cookie.Expires)
	// log.Println(cookie.Secure)
	// log.Println(cookie.HttpOnly)
	// log.Println(cookie.Path)
	// log.Println(cookie.SameSite)
	// log.Println(cookie.String())
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func (h *Handler) CorsWithRpID(c echo.Context) error {
	resp := map[string]string{"message": "OK"}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GroupApi(c echo.Context) error {
	resp := map[string]string{"text": "api"}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["userName"].(string)
	tokenType := claims["tokenType"].(string)
	if tokenType != "access" {
		resp := map[string]string{"message": "invalid or expired jwt!!"}
		return c.JSON(http.StatusUnauthorized, resp)
	}
	resp := map[string]string{"message": "Welcome " + name + "!"}
	return c.JSON(http.StatusOK, resp)
	// return c.JSON(http.StatusOK, claims)
}
