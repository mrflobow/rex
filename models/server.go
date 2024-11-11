package models

type Server struct {
	Host    string `yaml:"host"`
	KeyFile string `yaml:"key_file"`
	User    string `yaml:"user"`
}
