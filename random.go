package util

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 是否播随机种子, 用于解决多进程随机种子相同的情况
var seedFlag = false

// generate random string
func RandomStr(length int) string {
	if !seedFlag {
		// 如果没有播随机种子
		rand.Seed(time.Now().UnixNano())
		seedFlag = true
	}
	runes := make([]rune, length)
	lettersLen := len(letters)
	for i := 0; i < length; i++ {
		runes[i] = letters[rand.Intn(lettersLen)]
	}
	return string(runes)
}
