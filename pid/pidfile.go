package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Exists open a pid file.
func Exists(pidfile string) error {
	pid, err := readPidFile(pidfile)
	if err != nil {
		return fmt.Errorf("open pid file %s falied, %s", pidfile, err)
	}

	if !existsPID(pid) {
		return fmt.Errorf("not found process %v", pid)
	}

	return nil
}

// Open open a pid file and return this process.
func Open(pidfile string) (*os.Process, error) {
	pid, err := readPidFile(pidfile)
	if err != nil {
		return nil, fmt.Errorf("open pid file %s falied, %s", pidfile, err)
	}
	return os.FindProcess(pid)
}

// Generate create a pid file for own process.
func Generate(pidfile string, chDel <-chan struct{}) error {
	f, err := os.OpenFile(pidfile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintln(f, os.Getpid())
	if chDel != nil {
		go func() {
			select {
			case <-chDel:
				os.Remove(pidfile)
			}
		}()
	}
	return nil
}

// openPidFile read a pid file.
func readPidFile(filename string) (int, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(bs)))
	if err != nil {
		return 0, err
	}

	return n, nil
}

// existsPID check exists pid process.
func existsPID(pid int) bool {
	_, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	return true
}
