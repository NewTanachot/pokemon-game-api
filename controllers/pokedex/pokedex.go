package pokedexctr

import (
	pokedexusc "pokemon-game-api/usercases/pokedex"

	"github.com/gin-gonic/gin"
)

type IPokedexController interface {
	GetPokemonFromPokedex(c *gin.Context)
}

type PokedexController struct {
	PokedexUsecase pokedexusc.IPokedexUsecase
}
