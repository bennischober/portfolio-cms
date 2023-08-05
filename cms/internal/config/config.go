package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
    Database struct {
        Type string `yaml:"type" envconfig:"DB_TYPE"`
		Name string `yaml:"name" envconfig:"DB_NAME"`
        Connect string `yaml:"connect" envconfig:"DB_CONNECTION_STRING"`
    } `yaml:"database"`
    Server struct {
        Port string `yaml:"port" envconfig:"SERVER_PORT"`
    } `yaml:"server"`
	Logging struct {
		Level string `yaml:"level" envconfig:"LOG_LEVEL"`
		Port string `yaml:"port" envconfig:"LOG_PORT"`
	} `yaml:"logging"`
}


func LoadConfig(configPath string, logger *log.Logger) (*Config, error) {
	var cfg Config

	// Read from file
	f, err := os.Open(configPath)
	if err != nil {
		logger.Printf("Error opening config file: %v", err)
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.Printf("Error decoding config file: %v", err)
		return nil, err
	}

	// Read from environment variables
	err = envconfig.Process("", &cfg)
	if err != nil {
		logger.Printf("Error processing environment variables: %v", err)
		return nil, err
	}

	return &cfg, nil
}
