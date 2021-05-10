package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main(frag string) {

	if len(frag) <= 0 {
		return
	}

	// TODO: 二重がけ
	// TODO: パラメーター取得による複雑化
	result := sha256.Sum256([]byte(frag))

	// TODO: judge4digitsでの判定で再帰化
	fmt.Println(strings.ToUpper(hex.EncodeToString(result[:])))
}

func judge4digits(hash string){
	// TODO: 上4桁が0000かどうかの判定
}