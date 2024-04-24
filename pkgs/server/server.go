package server

import (
	"pokemon-game-api/pkgs/config"
	"pokemon-game-api/pkgs/di"
	customlog "pokemon-game-api/pkgs/logs"
	"pokemon-game-api/pkgs/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func GinSetup() {
	app = gin.Default()
	app.SetTrustedProxies([]string{})

	config.AddGodotEnvConfigurations()
	di.AddDependencyInjections()

	routes.MapControllerRouting(app)
}

func GinStart() {
	customlog.WriteInfoRuningServerPathLog(*config.Port)
	app.Run(*config.Port)
}
