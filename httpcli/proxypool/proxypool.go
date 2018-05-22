package proxypool

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/ckeyer/logrus"
)

const (
	ProxyPoolOptionCountry_domestic = "国内"
	ProxyPoolOptionCountry_foreign  = "国外"
)

var (
	DistantFuture = time.Hour * 24 * 365 * 100
)

type ProxyStatus uint8

// const (
// 	ProxyStatusOK          = iota
// 	ProxyStatusBanned      // 被临时封禁
// 	ProxyStatusUnreachable // 不可达
// )

var (
	ErrNotFound = errors.New("proxypool: not found a useful proxy url")
)

type ProxyPoolOption struct {
	PoolURL     string
	Transparent bool // 透明，否则高匿
	UseHttps    bool
	Count       int
	Country     string
	Area        string
}

func (opt *ProxyPoolOption) URL() *url.URL {
	pURL, err := url.Parse(opt.PoolURL)
	if err != nil {
		return nil
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

	if opt.Count > 0 {
		query.Set("count", strconv.Itoa(opt.Count))
	}

	if opt.Country != "" {
		query.Set("country", opt.Country)
	}
	if opt.Area != "" {
		query.Set("area", opt.Area)
	}

	pURL.RawQuery = query.Encode()

	return pURL
}

type ProxyPool struct {
	sync.Mutex

	opt    ProxyPoolOption
	pool   map[*url.URL]time.Time
	stopCh chan struct{}
}

func NewPool(opt ProxyPoolOption) (*ProxyPool, error) {
	p := &ProxyPool{
		opt:    opt,
		pool:   map[*url.URL]time.Time{},
		stopCh: make(chan struct{}),
	}

	err := p.updateProxyURLs()
	if err != nil {
		return nil, err
	}

	go p.runReplace()

	return p, nil
}

// runReplace
func (p *ProxyPool) runReplace() {
	for {
		select {
		case <-time.Tick(time.Minute * 25):
			err := p.updateProxyURLs()
			if err != nil {
				logrus.Errorf("load proxy pool failed, %s", err)
			}
		case <-p.stopCh:
			return
		}
	}
}

func (p *ProxyPool) updateProxyURLs() error {
	resp, err := http.Get(p.opt.URL().String())
	if err != nil {
		return err
	}

	var addrs [][]interface{}
	err = json.NewDecoder(resp.Body).Decode(&addrs)
	if err != nil {
		return err
	}

	scheme := "http"
	if p.opt.UseHttps {
		scheme = "https"
	}
	for _, addr := range addrs {
		if len(addr) != 3 {
			continue
		}
		u, err := url.Parse(fmt.Sprintf("%s://%s:%v/", scheme, addr[0], addr[1]))
		if err != nil {
			continue
		}
		p.add(u)
	}

	return nil
}

// Close
func (p *ProxyPool) Close() {
	select {
	case <-p.stopCh:
	default:
		close(p.stopCh)
	}
}

// Dirty
func (p *ProxyPool) Dirty(item *url.URL, wait time.Duration) {
	p.Lock()
	defer p.Unlock()
	if _, exists := p.pool[item]; exists {
		p.pool[item] = time.Now().Add(wait)
	}
}

// ProxyURL
func (p *ProxyPool) ProxyURL() (*url.URL, error) {
	p.Lock()
	defer p.Unlock()

	now := time.Now()
	for u, tm := range p.pool {
		if tm.Before(now) {
			return u, nil
		}
	}

	return nil, ErrNotFound
}

// ProxyClient
func (p *ProxyPool) NewClient() (*http.Client, *url.URL, error) {
	u, err := p.ProxyURL()
	if err != nil {
		return http.DefaultClient, nil, err
	}
	cli := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(u),
		},
		Timeout: time.Second * 10,
	}
	return cli, u, nil
}

// remove
func (p *ProxyPool) remove(item *url.URL) {
	p.Lock()
	defer p.Unlock()
	if _, exists := p.pool[item]; exists {
		p.pool[item] = time.Now().Add(DistantFuture)
	}
}

// add
func (p *ProxyPool) add(item *url.URL) {
	p.Lock()
	defer p.Unlock()
	if _, exists := p.pool[item]; !exists {
		p.pool[item] = time.Now()
	}
}
