package pokedexctr

import (
	"net/http"
	pokedexusc "pokemon-game-api/usercases/pokedex"

	"github.com/gin-gonic/gin"
)

type PokedexController struct {
	PokedexUsecase pokedexusc.IPokedexUsecase
}

func NewPokedexController(pokedexUsecase pokedexusc.IPokedexUsecase) IPokedexController {
	return PokedexController{PokedexUsecase: pokedexUsecase}
}

func (p PokedexController) GetPokemonFromPokedex(c *gin.Context) {
	result, cErr := p.PokedexUsecase.GetPokedex()

	if cErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, cErr)
		return
	}

	c.JSON(http.StatusOK, NewPokedexControllerResponse(result))
}
