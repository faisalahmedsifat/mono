package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Dir     string            `yaml:"dir"`
	Command string            `yaml:"command"`
	Watch   []string          `yaml:"watch"`
	Env     map[string]string `yaml:"env"`
	Port    int               `yaml:"port"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if len(cfg.Services) == 0 {
		return nil, fmt.Errorf("no services defined in config")
	}
	return &cfg, nil
}

func FilterServices(cfg *Config, onlyTasks []string) *Config {
	//TODO
	return cfg
}
