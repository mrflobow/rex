package services

import (
	"os"

	"github.com/mrflobow/rex/models"
	"gopkg.in/yaml.v3"
)

type ConfigLoader struct {
}

func (c *ConfigLoader) LoadConfig(config_file string) (*models.Config, error) {
	config := models.Config{}

	yamlData, err := os.ReadFile(config_file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(yamlData, &config); err != nil {
		return nil, err
	}

	return &config, nil

}
