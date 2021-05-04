package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API API `yaml:"api"`
	DB  DB  `yaml:"db"`
}

type API struct {
	APEX struct {
		Header string `yaml:"header"`
		Token  string `yaml:"token"`
	}
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func Load() (*Config, error) {
	// ===== Viper Settings
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("infrastructure/config/")

	// ===== Load
	if err := viper.ReadInConfig(); err != nil {
		var vError viper.ConfigFileNotFoundError
		ok := errors.As(err, &vError)

		if ok {
			return nil, fmt.Errorf("[Fatal] File Not Found: %w \n", err)
		} else {
			return nil, fmt.Errorf("[Fatal] File is exists, but other error happened: %w \n", err)
		}
	}

	// ===== Bind Configs to Struct
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("[Fatal] Unmarshal error: %w \n", err)
	}

	return &cfg, nil
}
