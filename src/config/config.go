package config

import (
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Server struct {
		Port uint32
	}
	Database struct {
		Url string
	}
	JWT struct {
		PrivateKeyPath string `mapstructure:"private_key_path"`
	}
}

var config GlobalConfig
var isLoaded = false

func GetConfig(forceReload ...bool) GlobalConfig {
	if !isLoaded || len(forceReload) > 0 && forceReload[0] {
		load()
	}
	return config
}

func load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	isLoaded = true
}
