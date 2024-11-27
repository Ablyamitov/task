package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Server `yaml:"server"`
	DB     `yaml:"db"`
}

type Server struct {
	Port int    `yaml:"port" env-default:"8080"`
	Host string `yaml:"host" env-default:"localhost"`
}

type DB struct {
	URL string `yaml:"url" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	config := Config{}
	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file is not exist: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("cannot read config: %s", err)
	}

	return &config, nil
}
