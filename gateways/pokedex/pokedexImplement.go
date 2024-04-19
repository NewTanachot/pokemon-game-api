package pokedexgwy

import (
	"encoding/json"
	"io"
	"net/http"
	"pokemon-game-api/domains/constants"
	"pokemon-game-api/pkgs/config"
	customerror "pokemon-game-api/pkgs/error"
)

type PokedexGateway struct{}

func NewPokedexGateway() IPokedexGateway {
	return PokedexGateway{}
}

func (p PokedexGateway) GetPokeapiPokedex() (*PokedexGatewayResponse, error) {
	response, err := http.Get(config.PokeapiBaseUrl + "pokedex/5")

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	pokeapiResponse := new(PokeapiPokedexResponse)

	if err = json.Unmarshal(responseBody, pokeapiResponse); err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.UnableToParseJsonToStruct)
	}

	result := PokedexGatewayResponse{
		Id:   uint(pokeapiResponse.ID),
		Name: pokeapiResponse.Name,
	}

	for _, v := range pokeapiResponse.Descriptions {
		if v.Language.Name == constants.En {
			result.Description = v.Description
		}
	}

	for _, v := range pokeapiResponse.PokemonEntries {
		result.Pokemons = append(result.Pokemons, PokedexPokemonDetail{
			Number: uint(v.EntryNumber),
			Name:   v.PokemonSpecies.Name,
			Url:    v.PokemonSpecies.URL,
		})
	}

	return &result, nil
}
