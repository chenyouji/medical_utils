package utils

import (
	"github.com/spf13/viper"
)

var config struct {
	Nacos Nacos
}

func InitViper(filepath string) {
	v := viper.New()
	v.SetConfigFile(filepath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_ = v.Unmarshal(&config)
}
