package redis

import (
	"fmt"
)

type Options struct {
	Host     string
	Port     int64
	Password string
	Url      string
	DB       int
}

func toUrl(o *Options) string {
	// redis://[password@]host[port][/db_num]
	password := o.Password
	if password != "" {
		password = password + "@"
	}
	return fmt.Sprintf("redis://%s%s:%d/0", password, o.Host, o.Port)
}

func NewOptions(host string, port int64, password string, db int) *Options {
	options := &Options{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
	options.Url = toUrl(options)
	return options
}
