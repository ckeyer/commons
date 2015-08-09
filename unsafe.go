package main

import (
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
