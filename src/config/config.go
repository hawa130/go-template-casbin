package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Server struct {
		Address string
	}
	Database struct {
		Url string
	}
	JWT struct {
		PrivateKeyPath string `mapstructure:"private_key_path"`
	}
	GraphQL struct {
		EndPoint           string `mapstructure:"endpoint"`
		Playground         bool   `mapstructure:"playground"`
		PlaygroundEndpoint string `mapstructure:"playground_endpoint"`
	} `mapstructure:"graphql"`
}

var config GlobalConfig
var isLoaded = false

var onConfigChange func()

func OnConfigChange(run func()) {
	onConfigChange = run
}

func GetConfig() GlobalConfig {
	if !isLoaded {
		initViper()
		load()
	}
	return config
}

func load() {
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	isLoaded = true
}

func initViper() {
	initDefault()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("toml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config changed:", e.Name)
		load()
		onConfigChange()
	})
}

func initDefault() {
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("graphql.playground", true)
	viper.SetDefault("graphql.playground_endpoint", "/playground")
	viper.SetDefault("graphql.endpoint", "/graphql")
}
