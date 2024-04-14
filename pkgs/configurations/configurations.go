package configurations

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once          sync.Once
	configuration ServerConfiguration
)

func NewSingletonServerConfiguration() ServerConfiguration {
	once.Do(func() {
		// load configuration from .env file
		godotenv.Load()

		configuration = ServerConfiguration{
			Port: getEnvPort(),
		}
	})

	return configuration
}

func getEnvPort() string {
	defaultPort := "3000"
	port := os.Getenv(port)

	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf(":%s", port)
}
