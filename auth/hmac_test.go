package auth

import (
	"testing"
)

func TestHMacSha1(t *testing.T) {
	key := "123"
	src := "123456"
	hash := "6c405d14046741981cf4cb8797fdbec6eff5edb2"

	thash := HMACSHA1([]byte(src), []byte(key))
	if hash != thash {
		t.Error(hash)
		t.Error(thash)
	}
}
