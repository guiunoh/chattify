package config

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Config(config interface{}, name string) {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	ext := filepath.Ext(name)
	v.SetConfigName(strings.TrimSuffix(name, ext))
	v.SetConfigType(ext[1:])

	if err := v.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "failed read config"))
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.Unmarshal(config); err != nil {
		panic(errors.Wrap(err, "failed read config"))
	}
}
