package container

import (
	"fmt"
	"log"
	"os"
	customlog "pokemon-game-api/pkgs/logs"
)

const (
	portConstant = "PORT"
)

type Configuration struct {
	Port string
}

var (
	Config Configuration
)

func AddConfigurations() {
	Config.setPortConfiguration()

	customlog.WriteBorderedInfoLog("Add Configuration to DI container")
}

func (c *Configuration) setPortConfiguration() {
	value := os.Getenv(portConstant)
	if value == "" {
		log.Fatal("can not get port value from .env file")
	}

	c.Port = fmt.Sprintf(":%s", value)
}
