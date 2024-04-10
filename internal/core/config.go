package core

import (
	"avito-tech-backend/internal/pkg/web"
	"avito-tech-backend/internal/storage"
	"github.com/spf13/viper"
)

type Config struct {
	Storage storage.Config   `yaml:"storage"`
	Server  web.ServerConfig `yaml:"server"`
}

func ParseConfig(loader *viper.Viper) (*Config, error) {
	cfg := &Config{}
	if err := loader.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := loader.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
