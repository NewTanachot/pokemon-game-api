package pokedexusc

import pokedexgwy "pokemon-game-api/gateways/pokedex"

type IPokedexUsecase interface {
	GetPokedex() (any, error)
}

type PokedexUsecase struct {
	PokedexGateway pokedexgwy.IPokedexGateway
}
