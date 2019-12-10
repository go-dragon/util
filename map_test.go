package util

import (
	"fmt"
	"testing"
)

// map functions

// get map keys
func TestGetMapKeys(t *testing.T) {
	var data = map[string]interface{}{
		"name":    "123",
		"age":     1,
		"profile": "good",
	}
	keys := GetMapKeys(data)
	fmt.Println(keys)
}

// get map keys
func TestGetMapKeys2(t *testing.T) {
	var data = map[string]string{
		"name":    "123",
		"profile": "good",
	}
	keys := GetMapKeys2(data)
	fmt.Println(keys)
}
