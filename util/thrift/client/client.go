package client

import (
	"math/rand"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Client struct {
	serverUris     []string
	serverUriIndex int
	Transport      *thrift.TBufferedTransport
	Protocol       *thrift.TBinaryProtocolFactory
}

const DefaultBufferSize = 1024 * 1024 * 5

func NewThriftClient(serverUris []string) *Client {
	c := &Client{
		serverUris:     serverUris,
		serverUriIndex: 0,
	}
	return c
}

func (c *Client) NextServer() {
	rand.Seed(time.Now().Unix())
	c.serverUriIndex = rand.Intn(len(c.serverUris))
}
func (c *Client) ServerUri() string {
	uri := c.serverUris[c.serverUriIndex]
	c.NextServer()
	return uri
}
