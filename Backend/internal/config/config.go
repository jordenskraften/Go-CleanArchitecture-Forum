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
	User     string `json:"user"`
	DbName   string `json:"dbname"`
	Password string `json:"password"`
	Host     int    `json:"host"`
	Port     int    `json:"port"`
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
