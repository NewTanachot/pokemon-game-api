package pokedexctr

import (
	"net/http"
	customerror "pokemon-game-api/pkgs/error"
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
	region := c.Query("region")
	result, cErr := p.PokedexUsecase.GetPokedex(region)

	if cErr != nil {
		pErr := customerror.ParseFrom(cErr)
		c.AbortWithStatusJSON(pErr.Status, pErr.GetError())
		return
	}

	c.JSON(http.StatusOK, NewPokedexControllerResponse(result))
}
