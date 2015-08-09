package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5Sum(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Sum(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
