package httpcli

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type CookieMap map[string]*http.Cookie

func (c CookieMap) GetValue(key string) string {
	if v, ok := c[key]; ok {
		return v.Value
	}
	return ""
}

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

type Client struct {
	*http.Client
	cookies CookieMap
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Jar: new(Jar),
		},
		cookies: make(CookieMap),
	}
}

func NewProxyClient(proxyURL string) (*Client, error) {
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: &http.Client{
			Jar: new(Jar),
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		},
		cookies: make(CookieMap),
	}, nil
}

func (cli *Client) IsExists(Url string) bool {
	resp, err := cli.Get(Url)
	if err != nil {
		return false
	}
	resp.Body.Close()
	if resp.StatusCode >= 400 {
		return false
	}
	return true
}

func (cli *Client) PostJSON(Url string, data interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}

	return cli.Post(Url, "application/json", buf)
}

func (cli *Client) PostFile(Url string, formName string, f *os.File) (*http.Response, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile(formName, f.Name())
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		return nil, err
	}

	w.Close()
	req, err := http.NewRequest("POST", Url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return cli.Do(req)
}
