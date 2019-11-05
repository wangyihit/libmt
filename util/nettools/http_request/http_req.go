package http_request

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	POST = 1
	GET  = iota
)

type HttpRequest struct {
	url       string
	data      []byte
	method    int
	headers   map[string]string
	proxyAddr string // proxy with schema
	expire    int64
}

func NewHttpRequest(url string, method int, headers map[string]string, data []byte, proxyAddr string, expire int64) (*HttpRequest, error) {
	h := &HttpRequest{
		url:       url,
		data:      data,
		method:    GET,
		proxyAddr: proxyAddr,
		headers:   headers,
		expire:    expire,
	}
	return h, nil
}

func (r *HttpRequest) SendRequest() ([]byte, error) {
	if r.expire != 0 {
		now := time.Now().Unix()
		if now > r.expire {
			return nil, errors.New("request expired")
		}
	}
	return SendRequest(r.url, r.method, r.data, r.headers, r.proxyAddr)
}

func SendRequest(requestUrl string, method int, data []byte, headers map[string]string, proxyAddr string) ([]byte, error) {
	var httpReq *http.Request
	var err error
	if method == POST {
		httpReq, err = http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(data)))
	} else if method == GET {
		httpReq, err = http.NewRequest("GET", requestUrl, nil)
	}

	if err != nil {
		return []byte("Create request failed"), err
	}
	if headers != nil {
		for k, v := range headers {
			httpReq.Header[k] = []string{v}
		}
	}
	client := &http.Client{}
	if proxyAddr != "" {
		if strings.Index(proxyAddr, "socks") > 0 {
		}
		proxyUrl, err := url.Parse(proxyAddr)
		if err != nil {
			return nil, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return []byte("Do request failed"), err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("status code is not 200")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
