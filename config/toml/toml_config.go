package toml

import (
	"fmt"
	"github.com/spf13/viper"
)

type TomlConfig struct {
	AppName string
	Log     LogConfig
}

// 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

var c TomlConfig

func init() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c)
}
func GetConfig() TomlConfig {
	return c
}
