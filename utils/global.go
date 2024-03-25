package utils

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var esMysqlConfig struct {
	Es    Elastic `json:"es"`
	Mysql MySQL   `json:"mysql"`
	Oss   Oss     `json:"oss"`
	Redis Redis   `json:"redis"`
}

func Inits(filepath string) {
	InitViper(filepath)
	content, err := InitNacos(&Config.Nacos)
	if err != nil {
		log.Fatal("获取nacos配置失败，错误为:", err)
	}
	_ = json.Unmarshal([]byte(content), &esMysqlConfig)
	InitEs(&esMysqlConfig.Es)
	Db = InitModel(&esMysqlConfig.Mysql)
	InitRedis(&esMysqlConfig.Redis)
}
