package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	JWTToken string `json:"jwt_token"`
}

func SaveToken(token string) error {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, ".democtlconfig")

	config := Config{JWTToken: token}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600) // 0600 ensures only the user can read/write
}

func LoadToken() (string, error) {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(configDir, ".democtlconfig")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	return config.JWTToken, nil
}
