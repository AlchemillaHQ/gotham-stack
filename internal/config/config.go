package config

import (
	"encoding/json"
	"os"
)

type Admin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Config struct {
	Environment   string  `json:"environment"`
	Port          int     `json:"port"`
	DatabaseURL   string  `json:"database_url"`
	LogLevel      string  `json:"log_level"`
	SessionSecret string  `json:"session_secret"`
	Admins        []Admin `json:"admins"`
	DataPath      string  `json:"data_path"`
}

var ParsedConfig *Config

func ParseConfig(path string) *Config {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	ParsedConfig = &Config{}
	err = decoder.Decode(ParsedConfig)

	if err != nil {
		panic(err)
	}

	err = SetupDataPath()

	if err != nil {
		panic(err)
	}

	return ParsedConfig
}

func SetupDataPath() error {
	if ParsedConfig.DataPath == "" {
		ParsedConfig.DataPath = "./data"
	}

	if _, err := os.Stat(ParsedConfig.DataPath); os.IsNotExist(err) {
		err := os.Mkdir(ParsedConfig.DataPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
