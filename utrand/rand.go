package utrand

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var (
	AlphaNumeric characterSet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789")
)

type characterSet []rune

func GenRandString(n int, chars characterSet) string {
	if chars == nil {
		chars = AlphaNumeric
	}
	var b strings.Builder
	for i := 0; i < n; i++ {
		randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b.WriteRune(chars[randInt.Int64()])
	}
	str := b.String()

	return str
}
