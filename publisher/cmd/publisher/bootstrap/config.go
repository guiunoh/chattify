package bootstrap

import (
	"flag"
	_mysql "publisher/infrastructure/mysql"
	_redis "publisher/infrastructure/redis"
	_config "publisher/pkg/config"
)

var Config struct {
	Service struct {
		Port int `yaml:"port"`
	} `yaml:"service"`
	DB  _mysql.Config `yaml:"db"`
	RDB _redis.Config `yaml:"rdb"`
}

func init() {
	var name string
	flag.StringVar(&name, "config", "./config/config.yaml", "config file name. --config=config.yaml")
	flag.Parse()
	_config.Config(&Config, name)
}
