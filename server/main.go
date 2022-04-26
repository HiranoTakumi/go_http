package main

import (
	"flag"
	"github.com/HiranoTakumi/server/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "server port")
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			req := c.Request()
			return req.URL.Path == "/healthcheck"
		},
		Format: `{` +
			`"app":"test",` +
			`"level":"info",` +
			`"msg":"",` +
			`"data":{"remote_ip":"${remote_ip}","method":"${method}","path":"${path}","status":${status},"latency":${latency},"latency_human":"${latency_human}","request_id":"${id}"},` +
			`"type":"access",` +
			`"time":"${time_rfc3339}"` +
			"}\n",
	}))

	allowedOrigins := []string{}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
	}))
	// e.IPExtractor = echo.ExtractIPDirect()
	e.IPExtractor = echo.ExtractIPFromXFFHeader(
	// echo.TrustLinkLocal(false),
	// echo.TrustPrivateNet(false),
	// echo.TrustLoopback(false),
	)
	//// Routes
	e.GET("/healthcheck", handlers.Healthcheck)
	e.GET("/", handlers.LateResponse)

	e.GET("/list", handlers.ListApi)
	e.GET("/api", handlers.GroupApi)
	e.GET("/cors", handlers.CorsWithRpID)
	e.POST("/login", handlers.Login)
	e.GET("getCookie", handlers.CookieApi)
	e.GET("validCookie", handlers.ValidCookie)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/welcome", handlers.Restricted)
	r.GET("/refresh", handlers.Refresh)
	r.GET("Api", handlers.Restricted)

	a := e.Group("/cors")
	a.Use(corsConfig)
	a.GET("/check", handlers.CorsWithRpID)

	//// Start server
	log.Println("port: " + *port)
	log.Fatal(e.Start(":" + *port))
}

func corsConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origins := map[string]string{
			"test_id": "http://localhost:8080",
		}
		req := c.Request()
		origin := req.Header.Get(echo.HeaderOrigin)

		rpID := req.Header.Get("X-RELYING-PARTY-ID")
		allowedOriginsByRpID := []string{origins[rpID]}
		for _, v := range allowedOriginsByRpID {
			if origin == v {
				return next(c)
			}
		}
		return c.NoContent(http.StatusNoContent)
	}
}
