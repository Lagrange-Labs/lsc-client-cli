package utils

import (
	"fmt"
	"os/exec"
)

// CheckDockerImageExists checks if a Docker image exists locally.
func CheckDockerImageExists(imageName string) error {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// PullDockerImage pulls a Docker image from the registry.
func PullDockerImage(imageName string) error {
	cmd := exec.Command("docker", "pull", imageName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// RunDockerImage runs a Docker image.
func RunDockerImage(imageName, configFileName string) error {
	cmd := exec.Command("docker", "run", "-v", fmt.Sprintf("./%s:/app/%s", configFileName, configFileName), imageName, "/app/lagrange-node", "run-client", "-c", fmt.Sprintf("/app/%s", configFileName))
	if err := cmd.Start(); err != nil {
		return err
	}

	// Detach from the container
	if err := cmd.Process.Release(); err != nil {
		return err
	}
	return nil
}
