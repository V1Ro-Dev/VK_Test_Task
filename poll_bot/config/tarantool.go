package config

import (
	"os"
)

const (
	backupURL  string = "tarantool:3301"
	backupUser string = "admin"
	backupPass string = "secret"
)

type TarantoolConfig struct {
	tarantoolURL  string
	tarantoolUser string
	tarantoolPass string
}

func NewTarantoolConfig() *TarantoolConfig {
	return &TarantoolConfig{
		tarantoolURL:  getUrlWithDefault("TARANTOOL_URL", backupURL),
		tarantoolUser: getUserWithDefault("TARANTOOL_USER", backupUser),
		tarantoolPass: getPassWithDefault("TARANTOOL_PASS", backupPass),
	}
}

func (p *TarantoolConfig) GetURL() string {
	return p.tarantoolURL
}

func (p *TarantoolConfig) GetUser() string {
	return p.tarantoolUser
}

func (p *TarantoolConfig) GetPass() string {
	return p.tarantoolPass
}

func getUrlWithDefault(name string, defaultVal string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return defaultVal
}

func getUserWithDefault(name string, defaultVal string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return defaultVal
}

func getPassWithDefault(name string, defaultVal string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return defaultVal
}
