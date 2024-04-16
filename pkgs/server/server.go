package server

import (
	"fmt"
	"log"
	"pokemon-game-api/pkgs/configurations"
	dependency "pokemon-game-api/pkgs/dependencies"
	"pokemon-game-api/pkgs/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func GinSetup() {
	app = gin.Default()
	app.SetTrustedProxies([]string{})

	if err := configurations.AddConfigurations(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := dependency.AddDependencyInjections(); err != nil {
		log.Fatalln(err.Error())
	}

	routes.UseControllerRouting(app)
}

func GinStart() {
	logRuningServerPath(configurations.Port)
	app.Run(configurations.Port)
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
