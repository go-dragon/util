package util

import (
	"math"
	"math/big"
	"regexp"
)

// 找出字符串中所有数字, 如果没有匹配的，则返回nil
func FindNumsInString(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}

// 截取2位小数，四舍五不入。 这里通过 ✖️100取math.Floor再➗100方式实现的两位小数
func TwoDecimalPlaces(value float64) float64 {
	floatRes := big.NewFloat(0)
	f1 := big.NewFloat(value)
	floatRes = floatRes.Add(f1, floatRes)
	floatRes = floatRes.Mul(big.NewFloat(100), floatRes)
	f, _ := floatRes.Float64()
	f = math.Floor(f)
	floatRes = big.NewFloat(f)
	floatRes = floatRes.Quo(floatRes, big.NewFloat(100))
	res, _ := floatRes.Float64()
	return res
}
