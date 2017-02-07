package fileutils

import (
	"testing"
)

func TestTar(t *testing.T) {
	err := Tar("./", "/tmp/tartest.tar")
	if err != nil {
		t.Error(err)
	}
}

func TestTgz(t *testing.T) {
	err := Tgz("./", "/tmp/tgztest.tar")
	if err != nil {
		t.Error(err)
	}
}
