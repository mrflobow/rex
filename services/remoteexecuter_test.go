package services

import (
	"testing"

	"github.com/mrflobow/rex/models"
)

func TestConnectionFailed(t *testing.T) {

	config := models.Config{
		Server: map[string]models.Server{
			"testserver": {Host: "192.168.0.2", KeyFile: "", User: "test"},
		},
	}

	args := []string{"ls"}
	rex := RemoteExecutor{}
	if _, err := rex.ExecuteCommand(&config, "testserver", args); err == nil {
		t.Fatal("Expected connection error")
	}

}
