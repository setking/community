package global

import (
	"myApp/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Trans        ut.Translator
	DB           *gorm.DB
	Redis        *redis.Client
	ServerConfig = &config.ServerConfig{}
	//NacosConfig  *config.NacosConfig = &config.NacosConfig{}
)
