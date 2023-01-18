package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}
type Mysql struct {
	Addr     string `yaml:"addr"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var configOnce = &sync.Once{}
var config = &Config{}

func Data() *Config {
	configOnce.Do(func() {
		b, err := os.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(b, config)
		if err != nil {
			panic(err)
		}
	})
	return config
}
