package utils

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/lagrange-node/logger"
)

const dockerComposeTemplate = `version: "3.7"

services:
  lagrange_client_{{.Network}}_{{.ChainName}}:
    container_name: lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}
    image: {{.DockerImage}}
    restart: always
    ports:
      - "{{.PrometheusPort}}:8080"
    environment:
      - LAGRANGE_NODE_CLIENT_BLSKEYSTOREPATH=/app/config/keystore/bls.key
      - LAGRANGE_NODE_CLIENT_BLSKEYSTOREPASSWORDPATH=/app/config/keystore/bls.pass
      - LAGRANGE_NODE_CLIENT_SIGNERECDSAKEYSTOREPATH=/app/config/keystore/signer.key
      - LAGRANGE_NODE_CLIENT_SIGNERECDSAKEYSTOREPASSWORDPATH=/app/config/keystore/signer.pass
    command:
      - "/bin/sh"
      - "-c"
      - "/app/lagrange-node run-client -c /app/config/client.toml"
    volumes:
      - {{.ConfigFilePath}}:/app/config/client.toml
      - {{.BLSKeystorePath}}:/app/config/keystore/bls.key
      - {{.BLSKeystorePasswordPath}}:/app/config/keystore/bls.pass
      - {{.SignerECDSAKeystorePath}}:/app/config/keystore/signer.key
      - {{.SignerECDSAKeystorePasswordPath}}:/app/config/keystore/signer.pass
      - lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}:$HOME/.lagrange
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"

volumes:
  lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}:
`

// CheckDockerImageExists checks if a Docker image exists locally.
func CheckDockerImageExists(imageName string) error {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	if err := cmd.Run(); err != nil {
		if err := pullDockerImage(imageName); err != nil {
			return fmt.Errorf("failed to pull docker image %s: %s", imageName, err)
		}
	}
	cmd = exec.Command("docker", "image", "inspect", imageName, "--format='{{index .RepoDigests 0}}'")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get docker image digest: %s", err)
	}
	logger.Infof("Docker image %s exists locally with digest: %s", imageName, output)

	return nil
}

func pullDockerImage(imageName string) error {
	cmd := exec.Command("docker", "pull", imageName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// GenerateDockerComposeFile generates a Docker Compose file.
func GenerateDockerComposeFile(imageName, prometheusPort, configFilePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %s", err)
	}
	workDir := filepath.Join(homeDir, ".lagrange")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create work directory: %s", err)
	}

	var dockerConfig config.DockerComposeConfig
	// Get chain name and bls pub key from config file path
	configFileName := filepath.Base(configFilePath)
	seps := strings.Split(configFileName, "_")
	dockerConfig.Network = seps[1]
	dockerConfig.ChainName = seps[2]
	dockerConfig.BLSPubKeyPrefix = strings.Split(seps[3], ".")[0]
	dockerConfig.DockerImage = imageName
	dockerConfig.ConfigFilePath = configFilePath
	dockerConfig.PrometheusPort = prometheusPort

	// Load the client config
	clientCfg, err := config.LoadClientConfig(configFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to load client config: %s", err)
	}
	dockerConfig.BLSKeystorePath = clientCfg.BLSKeystorePath
	dockerConfig.BLSKeystorePasswordPath = clientCfg.BLSKeystorePasswordPath
	dockerConfig.SignerECDSAKeystorePath = clientCfg.SignerECDSAKeystorePath
	dockerConfig.SignerECDSAKeystorePasswordPath = clientCfg.SignerECDSAKeystorePasswordPath

	tmpDocker, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		return "", err
	}
	dockerComposeFilePath := filepath.Join(workDir, fmt.Sprintf("docker-compose_%s_%s_%s.yml", dockerConfig.Network, dockerConfig.ChainName, dockerConfig.BLSPubKeyPrefix))
	dockerConfigFile, err := os.Create(dockerComposeFilePath)
	if err != nil {
		return "", err
	}
	defer dockerConfigFile.Close()
	if err := tmpDocker.Execute(dockerConfigFile, dockerConfig); err != nil {
		return "", err
	}

	logger.Infof("Generated Docker Compose file: %s", dockerComposeFilePath)
	return dockerComposeFilePath, nil
}

// RunDockerImage runs a Docker image.
func RunDockerImage(imageName, prometheusPort, configFilePath string) error {
	dockerComposeFilePath, err := GenerateDockerComposeFile(imageName, prometheusPort, configFilePath)
	if err != nil {
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
