package auth

import (
	"testing"
)

func TestHMacSha1(t *testing.T) {
	key := "123"
	src := "123456"
	hash := "6c405d14046741981cf4cb8797fdbec6eff5edb2"

	thash := HmacSha1([]byte(src), []byte(key))
	if hash != thash {
		t.Log(hash)
		t.Log(thash)
		t.Error("HMacSha1 test Failed...")
	}
}
