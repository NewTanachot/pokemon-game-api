package server

import (
	"fmt"
	"pokemon-game-api/pkgs/container"
	"pokemon-game-api/pkgs/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func GinSetup() {
	app = gin.Default()
	app.SetTrustedProxies([]string{})

	container.AddConfigurations()
	container.AddDependencyInjections()

	routes.UseControllerRouting(app)
}

func GinStart() {
	logRuningServerPath(container.Config.Port)
	app.Run(container.Config.Port)
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
