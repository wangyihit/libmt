package redis_options

import (
	"flag"
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
	return fmt.Sprintf("redis://$s%s:%d/0", password, o.Host, o.Port)
}

func NewOptions(host string, port int64, password string, db int, regArgs bool) *Options {
	options := &Options{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
	options.Url = toUrl(options)
	if regArgs == false {
		return options
	}
	flag.StringVar(&options.Host, "redis_host", host, "redis host ip")
	flag.Int64Var(&options.Port, "redis_port", port, "redis port")
	flag.StringVar(&options.Password, "redis_Password", password, "redis Password")
	flag.IntVar(&options.DB, "redis_db", db, "redis db number, 0-15")
	return options
}
