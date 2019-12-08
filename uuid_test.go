package util

import (
	"log"
	"testing"
)

// test generate uuid
func TestNewUUID(t *testing.T) {
	log.Println(NewUUID())
	log.Println(NewUUID())
	uuid, _ := NewUUID()
	log.Println(len(uuid))
}

func BenchmarkNewUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID()
	}
}
