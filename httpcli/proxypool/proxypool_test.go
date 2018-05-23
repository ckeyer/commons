package proxypool

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestFormatOption(t *testing.T) {
	opts := map[*ProxyPoolOption]string{
		&ProxyPoolOption{
			PoolURL:     "http://a.com",
			Transparent: true,
			UseHttps:    true,
			Count:       5,
		}: "http://a.com?count=5&protocol=1&types=1",
	}
	for opt, url := range opts {
		if opt.URL().String() != url {
			t.Errorf("opt: %+v !== %s", opt, url)
		}
	}
}

func TestProxyPool(t *testing.T) {
	u := os.Getenv("ProxyPoolURL")
	if u == "" {
		return
	}
	opt := ProxyPoolOption{
		PoolURL: u,
		Country: ProxyPoolOptionCountry_domestic,
	}
	pool, err := NewPool(opt)
	if err != nil {
		t.Error(err)
	}

	t.Logf("proxy pool ip count: %v", len(pool.pool))
	for i := 0; i < 10; i++ {
		pu, err := pool.ProxyURL()
		if err != nil {
			t.Error(err)
			return
		}
		cli := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(pu),
			},
			Timeout: time.Second * 5,
		}
		_, err = cli.Get("http://example.ckeyer.com/")
		if err != nil {
			t.Error(err)
		}
	}
}

// TestProxyClient
func TestProxyClient(t *testing.T) {

}
