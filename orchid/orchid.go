package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/moby/moby/container"
	"os"
)

func main() {
	//XXX: Include functions list , run (receive port range- OPTIONAL, image name), kill (image name)
	action := os.Args[1]
	switch action {
	case "list":
		list()
		//	case "run":
		//		run()
	case "kill":
		kill()
	default:
		fmt.Println("Arguments must be list, run or kill")
	}
}

func list() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.Ports)
	}
}

func run() {
	image := os.Args[2]
	ports := os.Args[3:4]
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	NewBaseContainer(id, root string)
	containers := cli.ContainerStart(context.Background(), image, ports)
}

func kill() {
	id := os.Args[2]
	signal := os.Args[3]

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers := cli.ContainerKill(context.Background(), id, signal)
	fmt.Println("Killing your container ... softly:\n")
	fmt.Println(containers)

}
