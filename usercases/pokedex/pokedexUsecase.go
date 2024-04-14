package pokedexusc

import pokedexgwy "pokemon-game-api/gateways/pokedex"

func NewPokedexUsecase(pokedexGateway pokedexgwy.IPokedexGateway) IPokedexUsecase {
	return PokedexUsecase{PokedexGateway: pokedexGateway}
}

func (p PokedexUsecase) GetPokedex() (any, error) {
	pokemons, err := p.PokedexGateway.GetPokeapiSinnohPokedex()

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}
