package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (Config, error) {

	var config Config
	vp := viper.New()

	vp.SetConfigName("env")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(path)
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println("err", err)
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, err
}

type Config struct {
	Slack SlackConfig `mapstructure:"slack"`
}
