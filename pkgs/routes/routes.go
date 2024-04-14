package routes

import (
	"log"
	"pokemon-game-api/pkgs/dependency"

	"github.com/gin-gonic/gin"
)

func UseControllerRouting(app *gin.Engine) {
	if !dependency.IsInjected() {
		log.Fatalln("invalid dependency")
	}

	apiRoute := app.Group("/api")
	apiVerRoute := apiRoute.Group("/v1")

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex", dependency.PokedexController.GetPokemonFromPokedex)
	pokedexRoute.GET("")
	pokedexRoute.GET("/:id")
}
