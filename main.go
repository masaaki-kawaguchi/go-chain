package main

import (
	"github.com/labstack/echo/v4"
	Log "github.com/sirupsen/logrus"
	"go-chain/server"
	"net/http"
)

func init() {
	Log.SetLevel(Log.DebugLevel)
	Log.SetFormatter(&Log.JSONFormatter{})
}

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

func main() {

	router := server.Init()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	router.POST("/mining", func(c echo.Context) error {
		// TODO: マイニングの処理に変える serverディレクトリに書く予定
		return c.String(http.StatusOK, "mining")
	})

	Log.Fatal(router.Start(":1323"))

}
