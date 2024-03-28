package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Nacos Nacos
}

func InitViper(filepath string, c *Config) {
	v := viper.New()
	v.SetConfigFile(filepath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_ = v.Unmarshal(&c)
}
