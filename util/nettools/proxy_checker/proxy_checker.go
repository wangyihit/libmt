package proxy_checker

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	HostCheckProxy      = 1
	HostCheckHttpServer = 1 << 1

	HostCheckStatusSuccess = 0
	HostCheckStatusFailed  = 1

	HostHasProxy      = 1
	HostHasHttpServer = 2

	ProxyTypeHttp   = 1
	ProxyTypeHttps  = 2
	proxyTypeSocks5 = 3
)

type HostChecker struct {
}

var DefaultCheckURL = "https://www.baidu.com"
var DefaultProxyIP = "109.161.48.228"
var DefaultProxyPort = 53281

func NewHostChechker() *HostChecker {
	c := &HostChecker{}
	return c
}

type ProxyInfo struct {
	Host string
	Port int
	Type int
}

func NewProxyHost(host string, port int, proxyType int) *ProxyInfo {
	i := &ProxyInfo{
		Host: host,
		Port: port,
		Type: proxyType,
	}
	return i
}

type HostCheckResult struct {
	Status       int
	ProxyStatus  int
	ProxyPort    int
	ProxyHost    string
	ServerStatus int
	Message      string
	ReqDelay     int64
}

func NewHostCheckResult() *HostCheckResult {
	h := &HostCheckResult{
		Status:       0,
		ProxyStatus:  0,
		ServerStatus: 0,
		ProxyPort:    0,
		ProxyHost:    "",
		Message:      "",
		ReqDelay:     0,
	}
	return h
}

func (c *HostChecker) BuildProxyUri(host string, port int, uriType int) (error, string) {
	var uri string
	switch uriType {
	case ProxyTypeHttp:
		uri = fmt.Sprintf("http://%s:%d", host, port)
		return nil, uri
	case proxyTypeSocks5:
		uri = fmt.Sprintf("http://%s:%d", host, port)
		return nil, uri
	default:
		return errors.New("uriType not support"), ""
	}

}

func (c *HostChecker) CheckProxy(proxyHost string, proxyPort int, proxyType int, checkURL string, timeout int) *HostCheckResult {
	hostCheckResult := NewHostCheckResult()
	err, proxyUri := c.BuildProxyUri(proxyHost, proxyPort, proxyType)
	hostCheckResult.ProxyHost = proxyHost
	hostCheckResult.ProxyPort = proxyPort
	if err != nil {
		hostCheckResult.Status = -1
		return hostCheckResult
	}
	proxyURL, err := url.Parse(proxyUri)
	if err != nil {
		hostCheckResult.Status = -2
		hostCheckResult.Message = fmt.Sprintf("msg=param_error, host=%s, port=%d", proxyHost, proxyPort)
		return hostCheckResult
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
	//generating the HTTP GET request
	request, err := http.NewRequest("GET", checkURL, nil)
	if err != nil {
		hostCheckResult.Status = -3
		hostCheckResult.Message = fmt.Sprintf("msg=build_requset_failed, host=%s, port=%d", proxyHost, proxyPort)
		return hostCheckResult
	}

	//calling the URL
	start := time.Now().UnixNano() / 1000000
	_, err = client.Do(request)
	hostCheckResult.Status = 0
	if err != nil {
		hostCheckResult.Message = fmt.Sprintf("msg=check_proxy_connect_failed, host=%s, port=%d", proxyHost, proxyPort)
		return hostCheckResult
	}
	end := time.Now().UnixNano() / 1000000
	hostCheckResult.ReqDelay = end - start
	hostCheckResult.ProxyStatus = 1

	return hostCheckResult
}
