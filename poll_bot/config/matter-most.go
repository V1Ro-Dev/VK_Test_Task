package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type MatterMost struct {
	MattermostUserName string `toml:"MM_USERNAME"`
	MattermostToken    string `toml:"MM_TOKEN"`
	MattermostServer   string `toml:"MM_SERVER"`
}

func LoadConfig(configPath string) (*MatterMost, error) {
	if configPath == "" {
		configPath = "/app/deploym/config/matter-most/config.toml"
	}

	var cfg MatterMost
	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load config in config.LoadConfig: %w", err)
	}
	return &cfg, nil
}
