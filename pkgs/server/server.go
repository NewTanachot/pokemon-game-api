package server

import (
	"fmt"
	"pokemon-game-api/pkgs/configurations"
	"pokemon-game-api/pkgs/dependency"
	"pokemon-game-api/pkgs/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func GinSetup() {
	app = gin.Default()
	app.SetTrustedProxies([]string{})

	dependency.UseDependencyInjection()
	routes.UseControllerRouting(app)
}

func GinStart() {
	config := configurations.NewSingletonServerConfiguration()

	logRuningServerPath(config.Port)
	app.Run(config.Port)
}

func logRuningServerPath(port string) {
	fmt.Println("")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("|                                            |")
	fmt.Printf("|   server runing on http://localhost%s   |\n", port)
	fmt.Println("|                                            |")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("")
}
