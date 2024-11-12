package models

type Config struct {
	Server           map[string]Server           `yaml:"server"`
	CommandTemplates map[string]CommandTemplates `yaml:"templates"`
}
