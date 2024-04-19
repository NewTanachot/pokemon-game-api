package pokedexctr

import "github.com/gin-gonic/gin"

type IPokedexController interface {
	GetPokemonFromPokedex(c *gin.Context)
}
