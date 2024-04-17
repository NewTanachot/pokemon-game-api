package routes

import (
	"pokemon-game-api/pkgs/container"

	"github.com/gin-gonic/gin"
)

func UseControllerRouting(app *gin.Engine) {
	apiRoute := app.Group("/api")
	apiVerRoute := apiRoute.Group("/v1")

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex", container.PokedexController.GetPokemonFromPokedex)
	pokedexRoute.GET("")
	pokedexRoute.GET("/:id")
}
