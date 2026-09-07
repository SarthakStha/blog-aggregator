package config

import (
	"encoding/json"
	"fmt"
	"os"
	// Why use path/filepath instead of path?
	// path is OS agnostic.
	// While filepath will adapt to the GOOS and GOARCH
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (userConfig *Config) SetUser(userName string) error {
	userConfig.CurrentUserName = userName
	return write(*userConfig)
}

func getConfigFilePath() (string, error) {
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDirPath, configFileName), nil
}

func Read() (Config, error) {
	configData := Config{}
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return configData, fmt.Errorf("Error: Unable to get home dir path, %w", err)
	}

	// Better to use Open() to prevent loading
	// very large volume of file data in memory.
	jsonData, err := os.ReadFile(configFilePath)
	if err != nil {
		return configData, fmt.Errorf("Error: Unable to read the file, %w", err)
	}

	if err := json.Unmarshal(jsonData, &configData); err != nil {
		return configData, fmt.Errorf("Error: Unable to parse the data, %w", err)
	}

	return configData, nil
}

func write(userConfig Config) error {
	data, err := json.Marshal(userConfig)
	if err != nil {
		return fmt.Errorf("Error: Unable to marshal the data, %w", data)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error: Unable to get home dir path, %w", err)
	}

	if err := os.WriteFile(configFilePath, data, 0666); err != nil {
		return fmt.Errorf("Error: Unable to write to file, %w", err)
	}

	return nil
}
