package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func Sha(frag string) {

	if len(frag) <= 0 {
		return
	}

	// TODO: 二重がけ
	// TODO: パラメーター取得による複雑化
	result := sha256.Sum256([]byte(frag))

	fmt.Println(strings.ToUpper(hex.EncodeToString(result[:])))
}