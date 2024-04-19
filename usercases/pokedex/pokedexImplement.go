package pokedexusc

import (
	"pokemon-game-api/domains/models"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	"strconv"
	"strings"
)

type PokedexUsecase struct {
	PokedexGateway pokedexgwy.IPokedexGateway
}

func NewPokedexUsecase(pokedexGateway pokedexgwy.IPokedexGateway) IPokedexUsecase {
	return PokedexUsecase{PokedexGateway: pokedexGateway}
}

func (p PokedexUsecase) GetPokedex() (*PokedexUsecaseResponse, error) {
	pokemons, cErr := p.PokedexGateway.GetPokeapiPokedex()

	if cErr != nil {
		return nil, cErr
	}

	result := PokedexUsecaseResponse{
		Id:          pokemons.Id,
		Name:        pokemons.Name,
		Description: pokemons.Description,
	}

	for _, v := range pokemons.Pokemons {
		result.Pokemons = append(result.Pokemons, models.PokedexPokemonDetail{
			Id:   getPokemonIdFromUrl(v.Url),
			Name: v.Name,
		})
	}

	return &result, nil
}

func getPokemonIdFromUrl(url string) uint {
	slice := strings.Split(url, "/")
	value := slice[len(slice)-2]
	intValue, err := strconv.Atoi(value)

	if err != nil {
		intValue = 0
	}

	return uint(intValue)
}
