package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorResponse(c echo.Context, code string) error {
	return c.JSON(http.StatusBadRequest, "{error: error}")
}