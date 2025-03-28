package config

import (
	"os"
)

const (
	backupURL string = "tarantool:3301"
)

type TarantoolConfig struct {
	tarantoolURL string
}

func NewTarantoolConfig() *TarantoolConfig {
	return &TarantoolConfig{
		tarantoolURL: getUrlWithDefault("TARANTOOL_URL", backupURL),
	}
}

func (p *TarantoolConfig) GetURL() string {
	return p.tarantoolURL
}

func getUrlWithDefault(name string, defaultVal string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return defaultVal
}
