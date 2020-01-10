package util

import (
	"fmt"
	"testing"
)

func TestRandomStr(t *testing.T) {
	fmt.Println(RandomStr(32))
}

// benchmark randomStr
func BenchmarkRandomStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomStr(16)
	}
}

func TestRandomNumber(t *testing.T) {
	fmt.Println(RandomNumber(6))
}

// benchmark randomStr
func BenchmarkRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNumber(16)
	}
}

func TestTrueRandomStr(t *testing.T) {
	fmt.Println(TrueRandomStr(6))
}

// benchmark randomStr
func BenchmarkTrueRandomStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrueRandomStr(16)
	}
}

func TestTrueRandomNumber(t *testing.T) {
	fmt.Println(TrueRandomNumber(6))
}

// benchmark randomStr
func BenchmarkTrueRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrueRandomNumber(16)
	}
}
