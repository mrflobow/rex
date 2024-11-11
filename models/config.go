package models

type Config struct {
	Server map[string]Server `yaml:"server"`
}
