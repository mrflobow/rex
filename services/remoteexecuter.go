package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/melbahja/goph"
	"github.com/mrflobow/rex/models"
)

type RemoteExecutor struct {
	Config *models.Config
}

type RemoteOutput struct {
	Data []byte
}

func (r *RemoteExecutor) ExecuteCommand(server string, args []string) (*RemoteOutput, error) {
	if r.Config == nil {
		errors.New("Config not initiated")
	}

	vserver, ok := r.Config.Server[server]
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
		parser := CommandParser{config: r.Config}
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

	rout := RemoteOutput{Data: out}

	return &rout, nil
}

func (r *RemoteExecutor) MultiExec(group string, args []string) error {

	c := make(chan *RemoteOutput)

	spawns := 0

	glist, ok := r.Config.Groups[group]

	if !ok {
		return errors.New("Group not found")
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
			fmt.Print(string(data.Data))
		}

	}

	return nil

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
