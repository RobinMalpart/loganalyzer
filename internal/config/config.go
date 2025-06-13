package config

import (
	"encoding/json"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}
