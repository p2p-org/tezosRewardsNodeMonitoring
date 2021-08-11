package checker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type containerChecker struct {
	name      string
	dockerCli *client.Client
	id        string
}

func (c *containerChecker) AssertRunning() (err error) {
	ctx := context.Background()
	resp, err := c.dockerCli.ContainerInspect(ctx, c.id)
	if err != nil {
		return err
	}
	if !resp.State.Running {
		return nil
	}
	return nil
}

func NewContainerChecker(containerName string, dockerCli *client.Client) (c Checker, err error) {
	ctx := context.Background()
	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, errors.New("Error fetching containers")
	}
	for _, cont := range containers {
		for _, name := range cont.Names {
			if name == containerName {
				c = &containerChecker{}
				c.(*containerChecker).id = cont.ID
				c.(*containerChecker).dockerCli = dockerCli
				c.(*containerChecker).name = containerName
				return c, nil
			}
		}
	}
	return nil, errors.New("ubable to find container with a given name")
}
