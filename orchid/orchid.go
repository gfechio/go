package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	image := os.Args[0]
	run := os.Args[1]

	if strings.Contains(image, "web") {
		port := 8001

		conn, err := net.Dial("tcp", "127.0.0.1:"+string(port))
		if err != nil {

		} else {
			port++
		}

		ctx := context.Background()
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}

		_, err = cli.ImagePull(ctx, "docker.io/library/"+image, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        image,
			Cmd:          []string{run},
			ExposedPorts: [string(port)],
		}, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
		}

		out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)

	} else {
		fmt.Println("We only support web aaplication so far")
	}
}
