package services

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/melbahja/goph"
	"github.com/mrflobow/rex/models"
)

type RemoteExecutor struct {
	config *models.Config
}

func NewRemoteExecutor(config *models.Config) *RemoteExecutor {
	return &RemoteExecutor{config: config}
}

func (r *RemoteExecutor) ExecuteCommand(server string, args []string) (*models.RemoteOutput, error) {
	if r.config == nil {
		errors.New("Config not initiated")
	}

	vserver, ok := r.config.Server[server]
	if !ok {
		errors.New("Server not found , update config")
	}

	key_path := r.replaceHomeDir(vserver.KeyFile)
	auth, err := goph.Key(key_path, "")

	if err != nil {
		return nil, err
	}

	client, err := goph.New(vserver.User, vserver.Host, auth)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	var cmd string

	if len(args) > 0 && strings.HasPrefix(args[0], ":") {
		prefix := args[0]
		template := prefix[1:]
		parser := NewCommandParser(r.config)
		subArgs := args[1:]
		cmd, err = parser.ParseCommand(template, subArgs)

		if err != nil {
			return nil, err
		}

	} else {
		cmd = strings.Join(args, " ")
	}

	out, err := client.Run(cmd)

	if err != nil {
		return nil, err
	}

	rout := models.RemoteOutput{Data: out, Server: server}

	return &rout, nil
}

func (r *RemoteExecutor) MultiExec(group string, args []string) (*[]models.RemoteOutput, error) {

	var output []models.RemoteOutput

	c := make(chan *models.RemoteOutput)

	spawns := 0

	glist, ok := r.config.Groups[group]

	if !ok {
		return nil, errors.New("Group not found")
	}

	for _, element := range glist {

		spawns++
		go func(r *RemoteExecutor, server string, args []string) {
			out, err := r.ExecuteCommand(server, args)

			if err != nil {
				log.Println(err)
				c <- nil
			} else {
				c <- out
			}

		}(r, element, args)
	}

	for i := 0; i < spawns; i++ {
		data := <-c
		if data != nil {
			output = append(output, *data)
		}

	}

	return &output, nil

}

func (r *RemoteExecutor) replaceHomeDir(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", home, 1)
	}

	return path
}
