package pid

import (
	"os"
	"path"
	"testing"
	"time"
)

// TestPID
func TestPID(t *testing.T) {
	pidfile := path.Join(os.TempDir(), "testpidfile")
	if err := Exists(pidfile); err == nil {
		t.Errorf("should return not found error")
	}

	chStop := make(chan struct{})
	if err := Generate(pidfile, chStop); err != nil {
		t.Errorf("generate pid file failed, %s", err)
	}
	defer os.Remove(pidfile)

	if err := Exists(pidfile); err != nil {
		t.Errorf("should return nil error")
	}
	chStop <- struct{}{}
	time.Sleep(time.Second)
	if err := Exists(pidfile); err == nil {
		t.Errorf("should return an error")
	}
}
