package response

import (
	"app/infra/errormessage"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SccessResponse(c echo.Context, res interface{}) error {
	return c.JSON(http.StatusOK, res)
}

func ErrorResponse(c echo.Context, code int) error {
	emap := errormessage.SettingError()

	if len(emap) == 0 {
		return c.JSON(http.StatusInternalServerError, "{message: システムエラー}")
	}

	return c.JSON(emap[code].Status, emap[code])
}