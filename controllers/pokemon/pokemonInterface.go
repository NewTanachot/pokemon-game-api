package pokemonctr

import "github.com/gin-gonic/gin"

type IPokemonController interface {
	GetPokemonDetailById(c *gin.Context)
}
