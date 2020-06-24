package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HeartbeatRoute returns a 200 okay always, used to test if server is running
func HeartbeatRoute(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
