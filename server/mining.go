package server

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

func Mining(BlockNumber int) echo.HandlerFunc{
	// TODO: genesisか判断する処理
	return func(c echo.Context) error {
		randomNonce()
		return c.JSON(http.StatusOK, BlockNumber)
	}
}

// 0~16の4乗の値をランダムで返す
func randomNonce() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(16 * 16 * 16 * 16)
}

// ハッシュの上4桁が0000かどうかの判定
func judge4digits(hash string) bool {
	return hash[:4] == "0000"
}