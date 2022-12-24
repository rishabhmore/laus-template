package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// The parsing struct for the json response
type response struct {
	Data string `json:"data"`
}

// Endpoint to check the status of the server. If it is up and running or not
func StatusCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Data: "Always hustling, Always Alive!🍲"})
}
