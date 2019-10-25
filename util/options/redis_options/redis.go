package redis_options
import (
	"flag"
)
type Options struct {
	Host string
	Port int64
	Password string
}

func NewOptions(host string, port int64, password string, regArgs bool) *Options {
	options := &Options{
		Host:host,
		Port:port,
		Password:password,
	}
	if regArgs == false {
		return options
	}
	flag.StringVar(&options.Host, "redis_host", host, "redis host ip")
	flag.Int64Var(&options.Port, "redis_port", port, "redis port")
	flag.StringVar(&options.Password, "redis_Password", password, "redis Password")
	return options
}


