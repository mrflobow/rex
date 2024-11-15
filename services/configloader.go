package services

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mrflobow/rex/models"
	"gopkg.in/yaml.v3"
)

type ConfigLoader struct {
}

func (c *ConfigLoader) LoadDefault() (*models.Config, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}
	config_file := filepath.Join(home, ".rex", "config.yml")
	return c.LoadConfig(config_file)
}

func (c *ConfigLoader) LoadConfig(config_file string) (*models.Config, error) {
	config := models.Config{}
	log.Printf("Loading config from %v", config_file)

	yamlData, err := os.ReadFile(config_file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(yamlData, &config); err != nil {
		return nil, err
	}

	return &config, nil

}
