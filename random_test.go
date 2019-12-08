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
