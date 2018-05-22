package proxypool

import (
	"encoding/json"
	"fmt"
	"io"
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

const (
	DefaultMaxRetryTimes = 3
	DefaultHTTPTimeout   = time.Second * 5
	DefaultWaitInterval  = time.Second * 15
)

var (
	DistantFuture = time.Hour * 24 * 365 * 100
)

// type ProxyStatus uint8

// const (
// 	ProxyStatusOK          = iota
// 	ProxyStatusBanned      // 被临时封禁
// 	ProxyStatusUnreachable // 不可达
// )

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
	pool   map[string]*proxyURL
	stopCh chan struct{}

	// client options
	pre          []PreHandler
	post         []PostHandler
	retryTimes   int
	waitInterval time.Duration
	timeout      time.Duration
}

type proxyURL struct {
	proxyURL  *url.URL
	scheduled time.Time
}

type PreHandler func(*http.Request) error

type PostHandler func(*http.Response, io.Reader) error

func NewPool(opt ProxyPoolOption) (*ProxyPool, error) {
	p := &ProxyPool{
		opt:          opt,
		pool:         map[string]*proxyURL{},
		stopCh:       make(chan struct{}),
		retryTimes:   DefaultMaxRetryTimes,
		timeout:      DefaultHTTPTimeout,
		waitInterval: DefaultWaitInterval,
	}

	err := p.refreshProxyURLs()
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
		case <-time.Tick(time.Minute * 10):
			err := p.refreshProxyURLs()
			if err != nil {
				logrus.Errorf("proxypool: load proxy pool failed, %s", err)
			}
		case <-p.stopCh:
			return
		}
	}
}

func (p *ProxyPool) refreshProxyURLs() error {
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
	if purl, exists := p.pool[item.String()]; exists {
		purl.scheduled = time.Now().Add(wait)
	}
}

// ProxyURL
func (p *ProxyPool) ProxyURL() (*url.URL, error) {
	p.Lock()
	defer p.Unlock()

	now := time.Now()
	for _, purl := range p.pool {
		if purl.scheduled.Before(now) {
			return purl.proxyURL, nil
		}
	}

	return nil, ErrNotFound
}

// remove
func (p *ProxyPool) remove(item *url.URL) {
	p.Lock()
	defer p.Unlock()
	if purl, exists := p.pool[item.String()]; exists {
		purl.scheduled = time.Now().Add(DistantFuture)
	}
}

// add
func (p *ProxyPool) add(item *url.URL) {
	p.Lock()
	defer p.Unlock()
	if _, exists := p.pool[item.String()]; !exists {
		p.pool[item.String()] = &proxyURL{proxyURL: item, scheduled: time.Now()}
	}
}

// sleep
func (p *ProxyPool) sleep() {
	time.Sleep(p.waitInterval)
}
