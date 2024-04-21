package config

import (
	"fmt"
	"os"
	"pokemon-game-api/pkgs/constants"
	customlog "pokemon-game-api/pkgs/logs"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once

	Port           string
	PokeapiBaseUrl string
)

func AddGodotEnvConfigurations() {
	once.Do(func() {
		// load configuration from .env file
		godotenv.Load()

		setPortConfiguration()
		setPokeapiBaseUrlConfiguration()

		customlog.WriteBorderedInfoLog("Add Configuration to DI container")
	})
}

func setPortConfiguration() {
	value := os.Getenv(constants.Port)
	if value == "" {
		customlog.WriteFatalSetGodotEnvFailLog("port")
	}

	Port = fmt.Sprintf(":%s", value)
}

func setPokeapiBaseUrlConfiguration() {
	PokeapiBaseUrl = os.Getenv(constants.PokeapiBaseUrl)

	if PokeapiBaseUrl == "" {
		customlog.WriteFatalSetGodotEnvFailLog("pokeapi base url")
	}
}
