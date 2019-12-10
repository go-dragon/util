package util

import (
	"fmt"
	"testing"
)

// map functions

// get map keys
func TestGetMapKeys(t *testing.T) {
	var data = map[string]interface{}{
		"name": "123",
		"age": 1,
		"profile": "good",
	}
	keys := GetMapKeys(data)
	fmt.Println(keys)
}
