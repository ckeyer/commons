package pid

import (
	"os"
	"path"
	"testing"
)

// TestPID
func TestPID(t *testing.T) {
	pidfile := path.Join(os.TempDir(), "testpidfile")
	if err := Exists(pidfile); err == nil {
		t.Errorf("should return not found error")
	}

	if err := Generate(pidfile); err != nil {
		t.Errorf("generate pid file failed, %s", err)
	}
	defer os.Remove(pidfile)

	if err := Exists(pidfile); err != nil {
		t.Errorf("should return nil error")
	}
}
