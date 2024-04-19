package routes

import (
	"pokemon-game-api/pkgs/di"

	"github.com/gin-gonic/gin"
)

func MapControllerRouting(app *gin.Engine) {
	apiRoute := app.Group("/api")
	apiVerRoute := apiRoute.Group("/v1")

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex")
	pokedexRoute.GET("", di.PokedexController.GetPokemonFromPokedex)
}
