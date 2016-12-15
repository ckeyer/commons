package httpclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Use IPProxys: (https://github.com/black-tech/IPProxyPool)
type Addr struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func (a *Addr) ToURL(isHttps bool) string {
	protocal := "http"
	if isHttps {
		protocal = "https"
	}
	return fmt.Sprintf("%s://%s:%v", protocal, a.IP, a.Port)
}

type ProxyOption struct {
	PoolURL     string
	Transparent bool // 透明，否则高匿
	UseHttps    bool
	Count       int
	Country     string
	Area        string
}

func (opt *ProxyOption) String() string {
	pURL, err := url.Parse(opt.PoolURL)
	if err != nil {
		return ""
	}
	query := url.Values{}

	if opt.Transparent {
		query.Set("types", "1")
	} else {
		query.Set("types", "0")
	}

	if opt.UseHttps {
		query.Set("protocol", "1")
	} else {
		query.Set("protocol", "0")
	}

	if opt.Count > 0 && opt.Count < 100 {
		query.Set("count", strconv.Itoa(opt.Count))
	} else {
		query.Set("count", "1")
	}

	if opt.Country != "" {
		query.Set("country", opt.Country)
	}
	if opt.Area != "" {
		query.Set("area", opt.Area)
	}

	pURL.RawQuery = query.Encode()

	return pURL.String()
}

func GetProxyURLs(opt *ProxyOption) ([]string, error) {
	cli := NewClient()

	poolUrl := opt.String()
	if poolUrl == "" {
		return nil, fmt.Errorf("invalid Proxy Option: %+v", opt)
	}

	resp, err := cli.Get(poolUrl)
	if err != nil {
		return nil, err
	}

	var addrs []Addr

	err = json.NewDecoder(resp.Body).Decode(&addrs)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, addr := range addrs {
		ret = append(ret, addr.ToURL(opt.UseHttps))
	}

	return ret, nil
}
