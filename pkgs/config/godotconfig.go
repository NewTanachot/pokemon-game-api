package config

import (
	"fmt"
	"os"
	customlog "pokemon-game-api/pkgs/logs"
	"pokemon-game-api/pkgs/utils/constants"

	"github.com/joho/godotenv"
)

var (
	Port           string
	PokeapiBaseUrl string
)

func AddGodotEnvConfigurations() {
	// load configuration from .env file
	godotenv.Load()

	setPortConfiguration()
	setPokeapiBaseUrlConfiguration()

	customlog.WriteBorderedInfoLog("Add Configuration to DI container")
}

func setPortConfiguration() {
	value := os.Getenv(constants.Port)
	if value == "" {
		customlog.WriteFatalSetGodotEnvFailLog("port")
	}

	Port = fmt.Sprintf(":%s", value)
}

func setPokeapiBaseUrlConfiguration() {
	value := os.Getenv(constants.PokeapiBaseUrl)
	if value == "" {
		customlog.WriteFatalSetGodotEnvFailLog("pokeapi base url")
	}

	PokeapiBaseUrl = fmt.Sprintf(":%s", value)
}
