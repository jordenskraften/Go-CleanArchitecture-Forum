package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB DatabaseConfig `json:"database_config"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	SSLmode  string `json:"sslmode"`
}

func ParseConfig(filePath string) (Config, error) {

	env := os.Getenv("environment")
	if env == "prod" {
		filePath = os.Getenv("config_path")
	}

	var data []byte
	var err error
	data, err = os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error on open file %v", err)
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("error on parse json %v", err)
	}
	return config, nil
}
