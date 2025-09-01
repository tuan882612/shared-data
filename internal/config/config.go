package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type GlobalConf struct {
	Config struct {
		Server ServerConf `yaml:"server"`
	} `yaml:"config"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewConfiguration(configPath string) (*GlobalConf, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	config := &GlobalConf{}
	if err := yaml.Unmarshal(data, config); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return config, nil
}
