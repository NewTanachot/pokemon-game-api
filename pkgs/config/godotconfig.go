package config

import (
	"fmt"
	"os"
	"pokemon-game-api/pkgs/constants"
	customlog "pokemon-game-api/pkgs/logs"
	"sync"

	"github.com/joho/godotenv"
)

// should restructure to singleton struct like database

var (
	once sync.Once

	Port           *string
	MongoHost      *string
	MongoPort      *string
	MongoUser      *string
	MongoPassword  *string
	PokeapiBaseUrl *string
)

func AddGodotEnvConfigurations() {
	once.Do(func() {
		// load configuration from .env file
		godotenv.Load()

		setApplicationPortConfiguration()
		setMongoDbConfiguration()
		setPokeapiBaseUrlConfiguration()

		customlog.WriteBorderedInfoLog("Add Configuration to DI container")
	})
}

func setApplicationPortConfiguration() {
	value := os.Getenv(constants.Port)
	if value == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.Port)
	}

	*Port = fmt.Sprintf(":%s", value)
}

func setPokeapiBaseUrlConfiguration() {
	*PokeapiBaseUrl = os.Getenv(constants.PokeapiBaseUrl)

	if *PokeapiBaseUrl == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.PokeapiBaseUrl)
	}
}

func setMongoDbConfiguration() {

	if *MongoHost = os.Getenv(constants.MongoHost); *MongoHost == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoHost)
	}

	if *MongoPort = os.Getenv(constants.MongoPort); *MongoPort == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoPort)
	}

	if *MongoUser = os.Getenv(constants.MongoUser); *MongoUser == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoUser)
	}

	if *MongoPassword = os.Getenv(constants.MongoPassword); *MongoPassword == "" {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoPassword)
	}
}
