package config

import (
	"fmt"
	"os"
	"pokemon-game-api/pkgs/constants"
	customlog "pokemon-game-api/pkgs/logs"
	stringutils "pokemon-game-api/pkgs/utils/string"
	"sync"

	"github.com/joho/godotenv"
)

// should restructure to singleton struct like database

var (
	once sync.Once
	// need to allocate pointer when init for access address of pointer after that.
	// ref https://edwinsiby.medium.com/runtime-error-invalid-memory-address-or-nil-pointer-dereference-golang-dd4a58ab7536
	Port           *string = new(string)
	MongoHost      *string = new(string)
	MongoPort      *string = new(string)
	MongoUser      *string = new(string)
	MongoPassword  *string = new(string)
	MongoDbName    *string = new(string)
	PokeapiBaseUrl *string = new(string)
	SecretKey      *string = new(string)
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
	if stringutils.IsNilOrEmpty(&value) {
		customlog.WriteGodotEnvFailPanicLog(constants.Port)
	}

	*Port = fmt.Sprintf(":%s", value)
}

func setPokeapiBaseUrlConfiguration() {
	*PokeapiBaseUrl = os.Getenv(constants.PokeapiBaseUrl)

	if stringutils.IsNilOrEmpty(PokeapiBaseUrl) {
		customlog.WriteGodotEnvFailPanicLog(constants.PokeapiBaseUrl)
	}
}

func setSecretKeyConfiguration() {

}

func setMongoDbConfiguration() {

	if *MongoHost = os.Getenv(constants.MongoHost); stringutils.IsNilOrEmpty(MongoHost) {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoHost)
	}

	if *MongoPort = os.Getenv(constants.MongoPort); stringutils.IsNilOrEmpty(MongoPort) {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoPort)
	}

	if *MongoUser = os.Getenv(constants.MongoUser); stringutils.IsNilOrEmpty(MongoUser) {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoUser)
	}

	if *MongoPassword = os.Getenv(constants.MongoPassword); stringutils.IsNilOrEmpty(MongoPassword) {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoPassword)
	}

	if *MongoDbName = os.Getenv(constants.MongoDbName); stringutils.IsNilOrEmpty(MongoDbName) {
		customlog.WriteGodotEnvFailPanicLog(constants.MongoDbName)
	}
}
