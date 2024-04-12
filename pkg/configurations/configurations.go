package configurations

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once          sync.Once
	configuration Configuration
)

func NewSingletonConfiguration() Configuration {
	once.Do(func() {
		// load configuration from .env file
		godotenv.Load()

		configuration = Configuration{
			Port: getEnvPort(),
		}
	})

	return configuration
}

func getEnvPort() string {
	defaultPort := "8080"
	port := os.Getenv(port)

	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf(":%s", port)
}
