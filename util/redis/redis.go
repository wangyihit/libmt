package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type FetchTask struct {
	SessionId  string            `json:"session_id"`
	Urls       map[string]string `json:"urls"`
	UserIP     string            `json:"user_ip"`
	TbUname    string            `json:"tb_uname"`
	TbUpass    string            `json:"tb_upass"`
	WaitQrcode int               `json:"wait_qrcode"`
}

type Client struct {
	Host    string
	Port    int
	Address string
}

func NewRedisClient(host string, port int) *Client {
	client := &Client{
		Host:    host,
		Port:    port,
		Address: fmt.Sprintf("%s:%d", host, port),
	}
	return client
}

func (r *Client) Get(key string) (string, error) {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return "", err
	}
	defer c.Close()
	c.Send("GET", key)
	c.Flush()
	data, err := redis.String(c.Receive()) // reply from SET
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *Client) Set(key string, value string) error {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return err
	}
	defer c.Close()
	c.Send("SET", key, value)
	c.Flush()
	redis.String(c.Receive()) // reply from HSET
	return nil
}

func (r *Client) SetEx(key string, expire int, value string) error {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return err
	}
	defer c.Close()
	expireStr := strconv.Itoa(expire)
	c.Send("SETEX", key, expireStr, value)
	c.Flush()
	redis.String(c.Receive()) // reply from HSET
	return nil
}
func (r *Client) GetWithRetry(key string, retry int, interval int) (string, error) {
	// retry
	// interval seconds
	var err error
	err = nil
	data := ""
	for ; 0 < retry; retry-- {

		data, err = r.Get(key)
		if err != nil {
			// has error, retry later
		} else {
			break
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return data, err
}

// if key not exist, return error
func (r *Client) HGet(key string, field string) (string, error) {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return "", err
	}
	defer c.Close()
	c.Send("HGET", key, field)
	c.Flush()
	data, err := redis.String(c.Receive()) // reply from HGET
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *Client) HSet(key string, field string, value string) error {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return err
	}
	defer c.Close()
	c.Send("HSET", key, field, value)
	c.Flush()
	redis.String(c.Receive()) // reply from HSET
	return nil
}

func (r *Client) HExists(key string, field string) (bool, error) {
	exists := false
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return exists, err
	}
	defer c.Close()
	c.Send("HEXISTS", key, field)
	c.Flush()
	d, err := redis.Int(c.Receive()) // reply from HEXISTS
	if err == nil && d == 1 {
		exists = true
	}
	return exists, err
}

// hmset [string]string
func (r *Client) HMSetSS(key string, data map[string]string) error {
	c, err := redis.Dial("tcp", r.Address)
	if err != nil {
		return err
	}
	defer c.Close()
	keyCount := len(data)
	buffer := make([]interface{}, keyCount*2+1)
	buffer[0] = key
	i := 1
	for k, v := range data {
		buffer[i] = k
		buffer[i+1] = v
		i += 2
	}
	c.Send("HMSET", buffer...)
	c.Flush()
	redis.String(c.Receive()) // reply from HMSET
	return nil
}
