package alarm

import (
	"io"
)

type Alarmer interface {
	Send(title string, level string, data []byte, args ...string) error
	SendTo(title string, level string, data []byte, toList ...string) (int, error)
}

type Content struct {
	io.Writer
	io.Reader
}
