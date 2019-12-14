package util

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var numbers = []rune("0123456789")

// 是否播随机种子, 用于解决多进程随机种子相同的情况
var seedFlag = false

// check rand seed
func checkSeed() {
	if !seedFlag {
		// 如果没有播随机种子
		rand.Seed(time.Now().UnixNano())
		seedFlag = true
	}
}

// generate random string
func RandomStr(length int) string {
	checkSeed()
	runes := make([]rune, length)
	lettersLen := len(letters)
	for i := 0; i < length; i++ {
		runes[i] = letters[rand.Intn(lettersLen)]
	}
	return string(runes)
}

// generate random number string
func RandomNumberStr(length int) string {
	checkSeed()
	runes := make([]rune, length)
	numbersLen := len(numbers)
	// 初始化第一个rune
	runes[0] = numbers[1+rand.Intn(numbersLen-1)]
	for i := 1; i < length; i++ {
		runes[i] = numbers[rand.Intn(numbersLen)]
	}
	return string(runes)
}
