package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/client"
)

func main() {
	// Create a new Docker client
	// client.NewClientWithOpts(client.FromEnv) uses DOCKER_HOST environment variable
	// for connection, which is compatible with Podman's Docker compatibility mode.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// Ping the Docker daemon to check if it's running and accessible
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cli.Ping(ctx)
	if err != nil {
		fmt.Printf("Docker daemon (or compatible service) is NOT running or accessible: %v\n", err)
		return
	}

	fmt.Println("Docker daemon (or compatible service) is running and accessible.")

	// Optionally, get and print the Docker version for more details
	version, err := cli.ServerVersion(ctx)
	if err != nil {
		fmt.Printf("Error getting Docker server version: %v\n", err)
		return
	}
	fmt.Printf("Server Version: %s\n", version.Version)
}