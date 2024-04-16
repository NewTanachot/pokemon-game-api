package configurations

import (
	"errors"
	"fmt"
	"os"
)

const (
	port = "PORT"
)

var (
	Port string
)

func AddConfigurations() error {
	// todo use cahnnel to return error
	port, err := getEnvPort()

	if err != nil {
		return err
	}

	Port = port

	return nil
}

func getEnvPort() (string, error) {
	port := os.Getenv(port)

	if port == "" {
		return "", errors.New("Can't get PORT from .env file")
	}

	return fmt.Sprintf(":%s", port), nil
}
