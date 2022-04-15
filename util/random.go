package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成一个随机数 在最大值最小值之间
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// 生成随机字符串 根据传入数字决定长度
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// 生成随机6位英文字母
func RandomOwner() string {
	return RandomString(6)
}

// 生成随机金额 0 - 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// 返回随机货币类型
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
