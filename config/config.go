package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	Token  string `json:"access_token"`
	Secret string `json:"secret_access_token"`
}

func GetConfig() (*Config, error) {
	var conf Config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Error: couldn't return the current user's home directory.")
	}

	configFile := homeDir + "/absamconfig.json"

	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.New("Error: couldn't find ~/absamconfig.json")
	}

	json.Unmarshal(raw, &conf)
	return &conf, nil
}
