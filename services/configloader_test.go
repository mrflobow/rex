package services

import (
	"errors"
	"io/fs"
	"testing"
)

func TestConfigNotFound(t *testing.T) {
	cl := ConfigLoader{}
	_, err := cl.LoadConfig("notfound.yml")

	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			t.Fatalf("Wrong error was thrown %v", err)
		}

	} else {
		t.Fatal("No error was thrown on file loading")
	}
}

func TestConfigLoad(t *testing.T) {
	cl := ConfigLoader{}
	config, err := cl.LoadConfig("../testdata/config_sample.yml")

	hostCheck := "testhost"

	if err != nil {
		t.Fatal("Config not found")
	}

	server, ok := config.Server["test1"]

	if !ok {
		t.Fatalf("Key was not found %v", "test1")
	}

	if server.Host != hostCheck {
		t.Fatalf("Host %v doesn't match with %v", server.Host, hostCheck)
	}

}

func TestCorruptConfigLoad(t *testing.T) {
	cl := ConfigLoader{}
	config, err := cl.LoadConfig("../testdata/config_corrupt.yml")
	if err != nil {
		t.Fatal("Config not found")
	}

	_, ok := config.Server["tokio"]
	if ok {
		t.Fatalf("Key was  found %v, expected to be nil", "tokio")
	}
}
