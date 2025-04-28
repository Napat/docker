package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", HealthCheck)

	// Start server
	if err := e.Start(":80"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("shutting down the server %s", err.Error())
	}
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
