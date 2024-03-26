package utils

import (
	"encoding/json"
	"log"
)

var EsMysqlConfig struct {
	Es    Elastic `json:"es"`
	Mysql MySQL   `json:"mysql"`
	Redis Redis   `json:"redis"`
}

func Inits(filepath string) {
	InitViper(filepath)
	content, err := InitNacos(&Config.Nacos)
	if err != nil {
		log.Fatal("获取nacos配置失败，错误为:", err)
	}
	_ = json.Unmarshal([]byte(content), &EsMysqlConfig)
	InitEs(&EsMysqlConfig.Es)
	Db = InitModel(&EsMysqlConfig.Mysql)
	InitRedis(&EsMysqlConfig.Redis)
}
