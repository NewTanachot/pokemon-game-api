package routes

import "github.com/gin-gonic/gin"

func UseControllerRouting(app *gin.Engine) {
	apiRoute := app.Group("/api")
	apiVerRoute := apiRoute.Group("/v1")

	// pokedex route
	pokedexRoute := apiVerRoute.Group("/pokedex")
	pokedexRoute.GET("")
	pokedexRoute.GET("/:id")
}
