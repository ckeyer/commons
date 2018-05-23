package proxypool

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/ckeyer/logrus"
)

var (
	ErrFirewallDeny = errors.New("proxypool: Firewall denied access")
	ErrTimeout      = errors.New("proxypool: Connection timeout")
	ErrIllegalBody  = errors.New("proxypool: Body contains illegal substr")
	ErrIllegalURL   = errors.New("proxypool: URL contains illegal substr")
	ErrNotFound     = errors.New("proxypool: Not found a useful proxy url")
)

// isTimeout
func IsTimeoutErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers")
}

// isConnectionErr
func IsConnectionErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "unexpected EOF")
}

func WithBodyContains(substr string) PostHandler {
	return func(_ *http.Response, body io.Reader) error {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)

		if strings.Contains(buf.String(), substr) {
			logrus.Debugf("body contains illegal substr: %s", substr)
			return ErrIllegalBody
		}
		return nil
	}
}

func WithURLContains(substr string) PostHandler {
	return func(resp *http.Response, _ io.Reader) error {
		if strings.Contains(resp.Request.URL.String(), substr) {
			logrus.Debugf("URL contains illegal substr: %s", substr)
			return ErrIllegalURL
		}
		return nil
	}
}
