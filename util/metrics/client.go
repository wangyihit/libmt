package metrics

import (
	"fmt"
	"net"
	"strconv"

	metrics_message "lib.mt/util/metrics/tmessage"

	"lib.mt/util/thrift/serializer"
)

type Client struct {
	host string
	port int
	addr string
}

func NewClient(host string, port int) *Client {
	c := &Client{
		host: host,
		port: port,
		addr: fmt.Sprintf("%s:%d", host, port),
	}
	return c
}

func (c *Client) EmitMessage(msg *metrics_message.TMessage) {
	data, err := serializer.ThriftObjectToBytes(msg)
	if err == nil {
		conn, err := net.Dial("udp", c.addr)
		defer conn.Close()
		if err != nil {
			return
		}
		conn.Write(data)
	}
}

func (c *Client) EmitTimer(t *Timer) {
	c.EmitTimerWithTags(t, "")
}

func (c *Client) EmitTimerWithPrefix(t *Timer, prefix string) {
	t.Name = fmt.Sprintf("%s.%s", prefix, t.Name)
	c.EmitTimer(t)
}
func (c *Client) EmitCounter(name string, count int64) {
	m := metrics_message.NewTMessage()
	m.MetricsType = metrics_message.MetricsType_Gauges
	m.Param = make(map[string]string)
	m.Name = name
	counter := strconv.FormatInt(count, 10)
	m.Param["count"] = counter
	c.EmitMessage(m)
}

func (c *Client) EmitCounterWithTags(name string, count int64, tags string) {
	m := metrics_message.NewTMessage()
	m.MetricsType = metrics_message.MetricsType_Counter
	m.Param = make(map[string]string)
	m.Name = name
	counter := strconv.FormatInt(count, 10)
	m.Param["count"] = counter
	m.Param["tags"] = tags
	c.EmitMessage(m)
}

func (c *Client) EmitGaugeWithTags(name string, count int64, tags string) {
	m := metrics_message.NewTMessage()
	m.MetricsType = metrics_message.MetricsType_Gauges
	m.Param = make(map[string]string)
	m.Name = name
	counter := strconv.FormatInt(count, 10)
	m.Param["count"] = counter
	m.Param["tags"] = tags
	c.EmitMessage(m)
}

func (c *Client) EmitTimerWithTags(t *Timer, tags string) {
	m := metrics_message.NewTMessage()
	m.MetricsType = metrics_message.MetricsType_Timer
	m.Param = make(map[string]string)
	m.Name = t.Name

	count := fmt.Sprintf("%d", t.End-t.Start)
	m.Param["count"] = count
	if tags != "" {
		m.Param["tags"] = tags
	}
	c.EmitMessage(m)
}

func (c *Client) EmitMCounterWithPrefix(name string, prefix string, count int64) {
	name = fmt.Sprintf("%s.%s", prefix, name)
	c.EmitCounter(name, count)
}
