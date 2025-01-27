package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config = Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, nil
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/" + configFileName, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	os.WriteFile(path, data, 0666)
	return nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	if err := write(*c); err != nil {
		return err
	}

	return nil
}
