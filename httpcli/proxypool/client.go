package proxypool

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/ckeyer/logrus"
)

// ProxyClient
func (p *ProxyPool) NewClient() (cli *http.Client, purl *url.URL, err error) {
	for i := 0; i < p.retryTimes; i++ {
		purl, err = p.ProxyURL()
		if err != nil {
			logrus.Warnf("proxypool: got proxy url failed at %v times, %s", i+1, err)
			p.sleep()
			continue
		}
		break
	}
	if err != nil {
		logrus.Errorf("proxypool: got proxy url failed.")
		return
	}

	cli = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(purl),
		},
		Timeout: p.timeout,
	}
	return
}

func (p *ProxyPool) WithTimeout(timeout time.Duration) *ProxyPool {
	p.Lock()
	defer p.Unlock()
	p.timeout = timeout
	return p
}

func (p *ProxyPool) WithWaitInterval(interval time.Duration) *ProxyPool {
	p.Lock()
	defer p.Unlock()
	p.waitInterval = interval
	return p
}

func (p *ProxyPool) WithRetryTimes(times int) *ProxyPool {
	p.Lock()
	defer p.Unlock()
	if times <= 1 {
		p.retryTimes = 1
	} else {
		p.retryTimes = times
	}
	return p
}

func (p *ProxyPool) WithPreHandle(hdls ...PreHandler) *ProxyPool {
	p.Lock()
	defer p.Unlock()
	p.pre = hdls
	return p
}

func (p *ProxyPool) WithPostHandle(hdls ...PostHandler) *ProxyPool {
	p.Lock()
	defer p.Unlock()
	p.post = hdls
	return p
}

// Do http.Do
func (p *ProxyPool) Do(req *http.Request) (resp *http.Response, err error) {
	for _, prefunc := range p.pre {
		if err := prefunc(req); err != nil {
			return nil, fmt.Errorf("proxypool preHandler: %s", err)
		}
	}

	var (
		cli  *http.Client
		purl *url.URL
	)
	for i := 0; i < p.retryTimes; i++ {
		cli, purl, err = p.NewClient()
		if err != nil {
			return nil, fmt.Errorf("proxypool: got http client falied, %s", err)
		}

		resp, err = cli.Do(req)
		if IsTimeoutErr(err) || IsConnectionErr(err) {
			logrus.Warnf("proxypool: do request at %v times failed, %s", i+1, err)
			p.Dirty(purl, time.Minute*30)
			i--
			continue
		}
		if err != nil {
			logrus.Warnf("proxypool: do request at %v times failed, %s", i+1, err)
			p.Dirty(purl, time.Minute*3)
			continue
		}

		if err = p.doPostHandlers(resp); err != nil {
			logrus.Warnf("proxypool.Do: got invalid response body at %v times failed, %s", i+1, err)
			i--
			p.Dirty(purl, time.Minute*30)
			continue
		}
		break
	}
	return
}

func (p *ProxyPool) doPostHandlers(resp *http.Response) error {
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	for _, postfunc := range p.post {
		if err := postfunc(resp, bytes.NewBuffer(bodyBytes)); err != nil {
			return err
		}
	}
	return nil
}
