package main

import (
	"github.com/labstack/echo/v4"
	Log "github.com/sirupsen/logrus"
	"go-chain/route"
	"net/http"
)

func init() {
	Log.SetLevel(Log.DebugLevel)
	Log.SetFormatter(&Log.JSONFormatter{})
}

func main() {

	router := route.Init()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	Log.Fatal(router.Start(":1323"))
}
