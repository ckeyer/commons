package httpclient

import (
	"testing"
)

func TestFormatOption(t *testing.T) {
	opts := map[*ProxyOption]string{
		&ProxyOption{
			PoolURL:     "http://a.com",
			Transparent: true,
			UseHttps:    true,
			Count:       5,
		}: "http://a.com?count=5&protocol=1&types=1",
	}
	for opt, url := range opts {
		if opt.String() != url {
			t.Errorf("opt: %+v !== %s", opt, url)
		}
	}
}
