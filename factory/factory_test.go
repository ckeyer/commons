package factory

import (
	"crypto/sha512"
	"strconv"
	"testing"
)

func hash(i int) {
	sha512.New().Sum([]byte("hello" + strconv.Itoa(i)))
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash(i)
	}
}
