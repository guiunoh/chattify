package bootstrap

import (
	_redis "consumer-forward/infrastructure/redis"
	_config "consumer-forward/pkg/config"
	"flag"
)

var Config struct {
	Service struct {
		Channel string `yaml:"channel"`
	} `yaml:"service"`
	Forward struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"forward"`
	RDB _redis.Config `yaml:"rdb"`
}

func init() {
	var name string
	flag.StringVar(&name, "config", "./config/config.yaml", "config file name. --config=config.yaml")
	flag.Parse()
	_config.Config(&Config, name)
}
