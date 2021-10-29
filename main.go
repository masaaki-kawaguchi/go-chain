package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	Log "github.com/sirupsen/logrus"
	"go-chain/presenter/http/router"
	"net/http"
)

func init() {
	Log.SetLevel(Log.DebugLevel)
	Log.SetFormatter(&Log.JSONFormatter{})
}

const helloMessage = "Hello, World!"

var BlockNumber = 1

type Transaction struct {
	SenderAddress string
	ReceiverAddress string
	Amount float32
	Memo string
	Significant string
}

type Block struct {
	BlockNumber int
	Nonce int
	Hash int
	PrevHash int
	Transaction
}

func main() *echo.Echo {

	router := router.Init()

	//router.GET("/", func(c echo.Context) error {
	//	// 仮置き
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//Log.Fatal(router.Start(":8080"))

	router.GET("/", helloHandler)

	return router
	// マイニング用
	//router.POST("/mining", router.Mining(BlockNumber))

	//Log.Fatal(router.Start(":8080"))

}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, helloMessage)
}
