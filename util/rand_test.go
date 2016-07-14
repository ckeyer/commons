package util

import (
	"fmt"
	"testing"
)

// TestRand
func TestRandBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomBytes(10))
	}
}

// TestRandString ...
func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomString(20))
	}
}
