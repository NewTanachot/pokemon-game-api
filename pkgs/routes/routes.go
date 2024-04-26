package routes

import (
	"pokemon-game-api/pkgs/di"
	"pokemon-game-api/pkgs/poc"
	stringutils "pokemon-game-api/pkgs/utils/string"

	"github.com/gin-gonic/gin"
)

func MapControllerRouting(app *gin.Engine) {
	apiRoute := app.Group("/api")
	apiVerRoute := apiRoute.Group("/v1")

	// auth
	authRoute := apiVerRoute.Group("/auth")
	authRoute.GET(stringutils.Empty, (*di.AuthController).GetAllUser)
	authRoute.GET(":id", (*di.AuthController).GetUserById)
	authRoute.POST(stringutils.Empty, (*di.AuthController).Register)

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex")
	pokedexRoute.GET(stringutils.Empty, (*di.PokedexController).GetPokedexDetail)

	// pokemon route
	pokemonRoute := apiVerRoute.Group("/pokemon")
	pokemonRoute.GET(":id", (*di.PokemonController).GetPokemonDetailById)

	// POC mongo db
	pocRoute := apiVerRoute.Group("/poc")
	pocRoute.GET("/mongo", poc.GetUser)
	pocRoute.POST("/mongo", poc.CreateUser)
}
