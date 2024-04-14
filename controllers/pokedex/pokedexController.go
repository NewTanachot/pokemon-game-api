package pokedexctr

import (
	"net/http"
	pokedexusc "pokemon-game-api/usercases/pokedex"

	"github.com/gin-gonic/gin"
)

func NewPokedexController(pokedexUsecase pokedexusc.IPokedexUsecase) IPokedexController {
	return PokedexController{PokedexUsecase: pokedexUsecase}
}

func (p PokedexController) GetPokemonFromPokedex(c *gin.Context) {
	pokemons, err := p.PokedexUsecase.GetPokedex()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, pokemons)
}
