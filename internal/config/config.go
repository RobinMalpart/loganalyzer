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
		return nil, err // Erreur d'ouverture du fichier
	}

	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, err // Erreur de parsing JSON
	}

	return configs, nil
}
