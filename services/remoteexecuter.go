package services

import (
	"errors"
	"strings"

	"github.com/melbahja/goph"
	"github.com/mrflobow/rex/models"
)

type RemoteExecutor struct {
}

type RemoteOutput struct {
	Data []byte
}

func (r *RemoteExecutor) ExecuteCommand(config *models.Config, server string, args []string) (*RemoteOutput, error) {

	vserver, ok := config.Server[server]
	if !ok {
		errors.New("Server not found , update config")
	}
	auth, err := goph.Key(vserver.KeyFile, "")

	if err != nil {
		return nil, err
	}

	client, err := goph.New(vserver.User, vserver.Host, auth)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	cmd := strings.Join(args, " ")

	out, err := client.Run(cmd)

	if err != nil {
		return nil, err
	}

	rout := RemoteOutput{Data: out}

	return &rout, nil
}
