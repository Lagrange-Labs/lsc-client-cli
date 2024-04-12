package utils

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/logger"
)

const dockerComposeTemplate = `version: "3.5"

services:
  client_{{.ChainName}}:
    container_name: {{.ChainName}}_{{.BLSPubKey}}
    image: {{.DockerImage}}
    restart: always
    command:
      - "/bin/sh"
      - "-c"
      - "/app/lagrange-node run-client -c /app/config/client.toml"
    volumes:
      - {{.ConfigFilePath}}:/app/config/client.toml
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
`

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
func RunDockerImage(imageName, configFilePath string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %s", err)
	}
	workDir := filepath.Join(homeDir, ".lagrange")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return fmt.Errorf("failed to create work directory: %s", err)
	}

	var dockerConfig config.DockerComposeConfig
	// Get chain name and bls pub key from config file path
	configFileName := filepath.Base(configFilePath)
	seps := strings.Split(configFileName, "_")
	dockerConfig.ChainName = seps[1]
	dockerConfig.BLSPubKey = strings.Split(seps[2], ".")[0]
	dockerConfig.DockerImage = imageName
	dockerConfig.ConfigFilePath = configFilePath

	tmpDocker, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		return err
	}
	dockerComposeFilePath := filepath.Join(workDir, fmt.Sprintf("docker-compose-%s-%s.yml", dockerConfig.ChainName, dockerConfig.BLSPubKey))
	dockerConfigFile, err := os.Create(dockerComposeFilePath)
	if err != nil {
		return err
	}
	defer dockerConfigFile.Close()
	if err := tmpDocker.Execute(dockerConfigFile, dockerConfig); err != nil {
		return err
	}

	cmd := exec.Command("docker", "compose", "-f", dockerComposeFilePath, "up", "-d")
	if err := cmd.Start(); err != nil {
		return err
	}
	logger.Infof("Started Docker container with docker-compose file: %s", dockerComposeFilePath)
	// Detach from the container
	if err := cmd.Process.Release(); err != nil {
		return err
	}
	return nil
}
