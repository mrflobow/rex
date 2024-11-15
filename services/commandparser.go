package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mrflobow/rex/models"
)

type CommandParser struct {
	config *models.Config
}

func NewCommandParser(config *models.Config) *CommandParser {
	return &CommandParser{config: config}
}

func (c *CommandParser) ParseCommand(template string, args []string) (string, error) {
	if c.config == nil {
		return "", errors.New("Config not initialized")
	}

	ct, ok := c.config.CommandTemplates[template]

	if !ok {
		return "", errors.New("Template not found")
	}

	parsed_data := ct.Command

	for index, element := range args {
		replace := fmt.Sprintf("${{%v}}", index)
		parsed_data = strings.Replace(parsed_data, replace, element, -1)
	}

	return parsed_data, nil

}
