package bootstrap

import (
	"flag"
	_redis "thumbnailer/infrastructure/redis"
	_config "thumbnailer/pkg/config"
)

var Config struct {
	Profile string `yaml:"profile"`
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
