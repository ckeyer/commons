package file

import (
	"os"
	"testing"
)

func TestTar(t *testing.T) {
	err := TarFile("./", "/tmp/tartest.tar")
	if err != nil {
		t.Error(err)
	}

	os.MkdirAll("/tmp/tar/out", 0755)

	err = UnTarFile("/tmp/tartest.tar", "/tmp/tar/out")
	if err != nil {
		t.Error(err)
	}
}

func TestTgz(t *testing.T) {
	err := TgzFile("./", "/tmp/tgztest.tar.gz")
	if err != nil {
		t.Error(err)
	}

	os.MkdirAll("/tmp/tgz/out", 0755)
	err = UnTgzFile("/tmp/tgztest.tar.gz", "/tmp/tgz/out")
	if err != nil {
		t.Error(err)
	}
}
