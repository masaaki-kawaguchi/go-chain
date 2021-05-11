package main

import (
	Log "github.com/sirupsen/logrus"
	"go-chain/server"
)

func init() {
	Log.SetLevel(Log.DebugLevel)
	Log.SetFormatter(&Log.JSONFormatter{})
}

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

func main() {

	router := server.Init()

	//router.GET("/", func(c echo.Context) error {
	//	// 仮置き
	//	return c.String(http.StatusOK, "Hello, World!")
	//})

	// マイニング用
	router.POST("/mining", server.Mining(BlockNumber))

	Log.Fatal(router.Start(":1323"))

}
