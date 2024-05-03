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
	authRoute.POST("/register", (*di.AuthController).Register)
	authRoute.POST("/login", (*di.AuthController).Login)

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex")
	pokedexRoute.GET(stringutils.Empty, (*di.PokedexController).GetPokedexDetail)

	// pokemon route
	pokemonRoute := apiVerRoute.Group("/pokemon")
	pokemonRoute.GET(":id", (*di.PokemonController).GetPokemonDetailById)

	// POC mongo db
	pocRoute := apiVerRoute.Group("/poc/mongo")
	pocRoute.GET("/user", poc.GetUser)
	pocRoute.GET("/user/pokemon", poc.GetUserWithPokemon)
	pocRoute.POST("/user", poc.CreateUser)
	pocRoute.PATCH("/user/pokemon/:id", poc.UpdateUserPokemon)
	pocRoute.PUT("/user/:id", poc.UpdateUser)
	pocRoute.DELETE("/user/:id", poc.DeleteUser)

	pocRoute.GET("/pokemon", poc.GetPokemon)
	pocRoute.POST("/pokemon", poc.CreatePokemon)

	pocRoute.DELETE("/user/collection", poc.DropUserCollection)
}
