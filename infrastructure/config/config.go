package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API API `yaml:"api"`
}

type API struct {
	APEX struct {
		Header string `yaml:"header"`
		Token  string `yaml:"token"`
	}
}

func Load() (*Config, error) {
	// ===== Viper Settings
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("infrastructure/config/")

	// ===== Load
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("[Fatal] File Not Found: %s \n", err)
		} else {
			return nil, fmt.Errorf("[Fatal] File is exists, but other error happened: %s \n", err)
		}
	}

	// ===== Bind Configs to Struct
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("[Fatal] Unmarshal error: %s \n", err)
	}

	return &cfg, nil
}
