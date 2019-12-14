package util

import (
	"fmt"
	"log"
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

func TestRandomNumberStr(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		numberStr := RandomNumberStr(6)
		if numberStr[0] == '0'  {
			log.Fatal("first letter cannot be `0`")
		}
	}
	fmt.Println(RandomNumberStr(6))
}


// benchmark randomStr
func BenchmarkRandomNumberStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNumberStr(16)
	}
}
