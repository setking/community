package initialize

import (
	"fmt"
	"myApp/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	path, _ := os.Getwd()
	pathFile := fmt.Sprintf("%s/config.yaml", path)
	v := viper.New()
	v.SetConfigFile(pathFile)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(global.ServerConfig)
	if err != nil {
		panic(err.Error())
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改...")
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			zap.S().Error("viper.Unmarshal failed", zap.Error(err))
		}
	})
}
