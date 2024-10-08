package base62

import (
	"math/big"
)

const (
	base         = 62
	characterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var base62CharsetMap = map[rune]int{}

func init() {
	for i, c := range characterSet {
		base62CharsetMap[c] = i
	}
}

func Encode(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	charset := characterSet

	b = append([]byte{1}, b...)
	num := new(big.Int).SetBytes(b)
	chars := []rune{}
	zero := big.NewInt(0)
	for num.Cmp(zero) > 0 {
		var remainder *big.Int
		num, remainder = new(big.Int).DivMod(num, big.NewInt(base), big.NewInt(base))
		chars = append(chars, rune(charset[int(remainder.Int64())]))
	}
	if len(chars) == 0 {
		return "0"
	}
	return string(chars)
}

func Decode(s string) []byte {
	m := base62CharsetMap
	val := new(big.Int)
	baseMul := new(big.Int).SetInt64(1)
	for _, c := range s {
		remainder := new(big.Int).SetInt64(int64(m[c]))
		val.Add(val, new(big.Int).Mul(remainder, baseMul))
		baseMul.Mul(baseMul, big.NewInt(base))
	}
	bytes := val.Bytes()
	if len(bytes) == 0 {
		return []byte{}
	}
	return bytes[1:]
}
