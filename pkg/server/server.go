package server

import (
	"fmt"
	"pokemon-game-api/pkg/configurations"
	"pokemon-game-api/pkg/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func GinSetup() {
	app = gin.New()
	app.SetTrustedProxies([]string{})

	routes.UseControllerRouting(app)
}

func GinStart() {
	config := configurations.NewSingletonConfiguration()

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
