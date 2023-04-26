package bootstrap

import (
	"flag"
	_redis "notifier/infrastructure/redis"
	_config "notifier/pkg/config"
)

var Config struct {
	Service struct {
		Port int `yaml:"port"`
	} `yaml:"service"`
	RDB _redis.Config `yaml:"rdb"`
}

func init() {
	var name string
	flag.StringVar(&name, "config", "./config/config.yaml", "config file name. --config=config.yaml")
	flag.Parse()
	_config.Config(&Config, name)
}
