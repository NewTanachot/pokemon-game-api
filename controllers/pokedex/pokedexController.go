package pokedexctr

import "pokemon-game-api/usercases/pokedex"

type IPokedexController interface {
}

type PokedexController struct {
	PokedexUsecase pokedexusc.IPokedexUsecase
}

func NewPokedexController(pokedexUsecase pokedexusc.IPokedexUsecase) IPokedexController {
	return PokedexController{PokedexUsecase: pokedexUsecase}
}
