package services

import (
	"testing"

	"github.com/mrflobow/rex/models"
)

func loadConfig(t *testing.T) (*models.Config, error) {
	c := ConfigLoader{}
	config, err := c.LoadConfig("../testdata/config.yml")

	if err != nil {
		t.Fatal("Cannot load config")
	}
	return config, nil
}
func TestParsing(t *testing.T) {
	config, _ := loadConfig(t)

	cp := CommandParser{config: config}

	args := []string{"is", "a", "super"}
	out, err := cp.ParseCommand("hello_world", args)

	if err != nil {
		t.Fatal(err)
	}

	if out != "echo \"This is a super test\"" {
		t.Log(out)
		t.Fatal("Command string doesn't match")
	}
}
