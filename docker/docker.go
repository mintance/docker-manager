package docker

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/Sirupsen/logrus"
)

func getRunningContainers() ([]types.Container, error) {

	cli, err := client.NewEnvClient()

	if err != nil {

		logrus.Error("Can`t connect to Docker: " + err.Error())

		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {

		logrus.Error("Can`t get running containers from Docker: " + err.Error())

		return nil, err
	}

	for key, container := range containers {
		containers[key].ID = container.ID[:10]
		//fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return containers, nil
}

//func startContainer() (types.Container, error) {}
//
//func stopContainer() error {}
//
//func imagePull() {}