package pokedexgwy

import (
	"encoding/json"
	"io"
	"net/http"
	"pokemon-game-api/pkgs/config"
	"pokemon-game-api/pkgs/utils/constants"
)

type IPokedexGateway interface {
	GetPokeapiSinnohPokedex() (*pokedexGatewayResponse, error)
}

type PokedexGateway struct{}

func NewPokedexGateway() IPokedexGateway {
	return PokedexGateway{}
}

func (p PokedexGateway) GetPokeapiSinnohPokedex() (*pokedexGatewayResponse, error) {
	response, err := http.Get(config.PokeapiBaseUrl + "pokedex/5")

	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	pokeapiResponse := new(pokeapiPokedexResponse)

	if err = json.Unmarshal(responseBody, pokeapiResponse); err != nil {
		return nil, err
	}

	result := pokedexGatewayResponse{
		Id:   uint(pokeapiResponse.ID),
		Name: pokeapiResponse.Name,
	}

	for _, v := range pokeapiResponse.Descriptions {
		if v.Language.Name == constants.En {
			result.Description = v.Description
		}
	}

	for _, v := range pokeapiResponse.PokemonEntries {
		result.Pokemons = append(result.Pokemons, pokedexPokemonDetail{
			Number: uint(v.EntryNumber),
			Name:   v.PokemonSpecies.Name,
			Url:    v.PokemonSpecies.URL,
		})
	}

	return &result, nil
}
