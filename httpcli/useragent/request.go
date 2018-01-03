package useragent

import (
	"io"
	"net/http"
)

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", RandUserAgent())
	return req, nil
}
