package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/proxy"
)

const (
	// PROXY_ADDR = "n.ckeyer.com:22071"
	PROXY_ADDR = "127.0.0.1:1080"
	URL        = "http://d.ckeyer.com/"
)

// TestSock5
func TestSock5(t *testing.T) {
	return
	auth := &proxy.Auth{
		Password: "sXOLs3di8qgGB4tRDGUUNKlGD+Amgw",
	}

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, auth, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	httpClient := &http.Client{
		// setup a http client
		Transport: &http.Transport{
			// set our socks5 as the dialer
			Dial: dialer.Dial,
		},
	}

	// create a request
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("User-Agent", "hahhaa")

	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		os.Exit(3)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		os.Exit(4)
	}
	fmt.Println(string(b))
}
