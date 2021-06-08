package router

import (
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	return e
}