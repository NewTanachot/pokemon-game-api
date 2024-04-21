package pokedexusc

import (
	"pokemon-game-api/domains/models"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	"pokemon-game-api/pkgs/constants"
	"pokemon-game-api/pkgs/utils"
	"strconv"
	"strings"
)

type PokedexUsecase struct {
	PokedexGateway pokedexgwy.IPokedexGateway
}

func NewPokedexUsecase(pokedexGateway pokedexgwy.IPokedexGateway) IPokedexUsecase {
	return PokedexUsecase{PokedexGateway: pokedexGateway}
}

func (p PokedexUsecase) GetPokedex(region string) (*PokedexUsecaseResponse, error) {
	regionNo := utils.GetRegionNo(region)
	gwyResponse, cErr := p.PokedexGateway.GetPokeapiPokedex(regionNo)

	if cErr != nil {
		return nil, cErr
	}

	result := PokedexUsecaseResponse{
		Id:   gwyResponse.ID,
		Name: gwyResponse.Name,
	}

	for _, v := range gwyResponse.Descriptions {
		if v.Language.Name == constants.En {
			result.Description = v.Description
		}
	}

	for _, v := range gwyResponse.PokemonEntries {
		result.Pokemons = append(result.Pokemons, models.PokedexPokemon{
			Id:   getPokemonIdFromUrl(v.PokemonSpecies.URL),
			Name: v.PokemonSpecies.Name,
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
