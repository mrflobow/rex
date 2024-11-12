package services

import (
	"log"
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
	rex := RemoteExecutor{Config: &config}
	if _, err := rex.ExecuteCommand("testserver", args); err == nil {
		t.Fatal("Expected connection error")
	}

}

func TestSingleCommandExec(t *testing.T) {
	var load ConfigLoader
	config, err := load.LoadConfig("../testdata/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	rex := RemoteExecutor{Config: config}
	args := []string{"echo \"hello world\""}

	out, err := rex.ExecuteCommand("test1", args)

	if err != nil {
		log.Fatal(err)
	}

	sout := string(out.Data)
	if sout != "hello world\n" {
		log.Fatalf("Result mismatches %v , expected: hello world", sout)
	}
}
