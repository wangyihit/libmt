package nettools

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"errors"
	"time"
)

const (
	HostCheckProxy = 1
	HostCheckHttpServer = 1 << 1

	HostCheckStatusSuccess = 0
	HostCheckStatusFailed = 1

	HostHasProxy = 1
	HostHasHttpServer = 2

	ProxyTypeHttp = 1
	ProxyTypeHttps = 2
	proxyTypeSocks5 = 3
)
type HostChecker struct {

}

func NewHostChechker() *HostChecker {
	c := &HostChecker{}
	return c
}


type HostCheckResult struct {
	Status      int
	ProxyStatus int
	ProxyPort   int
	ServerStatus      int
	Message     string
}

func NewHostCheckResult()* HostCheckResult {
	h = & HostCheckResult{
		Status:      0,
		ProxyStatus: -1,
		ServerStatus: -1,
		ProxyPort:   0,
		Message:     "",

	}
}
var DefaultCheckURL = "http://120.92.79.94/api/feed/"
var proxyIP = "109.161.48.228"
var proxyPort = 53281


func (c *HostChecker) CheckIP(host string, port int){

}
func (c *HostChecker) BuildProxyUri(host string, port int, uriType int) (error, string){
	var uri string
	switch uriType {
	case ProxyTypeHttp:
		uri = fmt.Sprintf("http://%s:%d", host, port)
		return nil ,uri
	case proxyTypeSocks5:
		uri = fmt.Sprintf("http://%s:%d", host, port)
		return nil, uri
	default:
		return errors.New("uriType not support"), ""
	}

}
func (c *HostChecker)CheckProxy(proxyHost string, proxyPort int, proxyType int, checkURL string, timeout int) *HostCheckResult {
	hostCheckResult := NewHostCheckResult()
	err, proxyUri := c.BuildProxyUri(proxyHost, proxyPort, proxyType)
	if err != nil {
		hostCheckResult.Status = -1
		return hostCheckResult
	}
	proxyURL, err := url.Parse(proxyUri)
	if err != nil {
		hostCheckResult.Status = -1
		hostCheckResult.Message = fmt.Sprintf("msg=param_error, host=%s, port=%d", proxyHost, proxyPort)
		return hostCheckResult
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(timeout) * time.Second,
	}
	//generating the HTTP GET request
	request, err := http.NewRequest("GET", checkURL, nil)
	if err != nil {

	}

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.Status)
	return hostCheckResult
}