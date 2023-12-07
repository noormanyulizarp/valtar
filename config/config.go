// valtar/config/config.go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TwitterApiKey            string `json:"TwitterApiKey"`
	TwitterApiSecret         string `json:"TwitterApiSecret"`
	TwitterAccessToken       string `json:"TwitterAccessToken"`
	TwitterAccessTokenSecret string `json:"TwitterAccessTokenSecret"`
}

func LoadConfig(filename string) (*Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
