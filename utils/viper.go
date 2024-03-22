package utils

import (
	"github.com/spf13/viper"
)

var config struct {
	Nacos Nacos
}

func InitViper() {
	v := viper.New()
	v.SetConfigFile("./etc/goods.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_ = v.Unmarshal(&config)
}
