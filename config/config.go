package config

import (
	"encoding/json"
	"os"
)

type ConnConfig struct {
	DBHost     string `json:"dbhost"`
	DBPort     string `json:"dbport"`
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	Database   string `json:"database"`
}

type Config struct {
	Gobel ConnConfig `json:"gobel"`
	Rubel ConnConfig `json:"rubel"`
}

func NewConfig() *Config {
	config := &Config{}
	b, _ := os.ReadFile("env.json")
	json.Unmarshal(b, config)
	return config
}
