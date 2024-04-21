package pokemonctr

import (
	"net/http"
	customerror "pokemon-game-api/pkgs/error"
	pokemonusc "pokemon-game-api/usercases/pokemon"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	PokemonUsecase pokemonusc.IPokemonUsecase
}

func NewPokemonControllor(pokemonUsecase pokemonusc.IPokemonUsecase) IPokemonController {
	return PokemonController{PokemonUsecase: pokemonUsecase}
}

func (p PokemonController) GetPokemonDetailById(c *gin.Context) {
	id := c.Param("id")
	uscResponse, cErr := p.PokemonUsecase.GetPokemonDetailById(id)

	if cErr != nil {
		pErr := customerror.ParseFrom(cErr)
		c.AbortWithStatusJSON(pErr.Status, pErr.GetError())
		return
	}

	result := PokemonControllerResponse{
		Id:        uscResponse.Id,
		Name:      uscResponse.Name,
		Height:    uscResponse.Height,
		Weight:    uscResponse.Weight,
		Types:     uscResponse.Types,
		Abilities: uscResponse.Abilities,
		Stats:     uscResponse.Stats,
		Sound:     uscResponse.Sound,
		Avatar:    uscResponse.Avatar,
	}

	c.JSON(http.StatusOK, result)
}
