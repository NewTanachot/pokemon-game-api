package pokemongwy

import (
	"encoding/json"
	"io"
	"net/http"
	"pokemon-game-api/pkgs/config"
	"pokemon-game-api/pkgs/constants"
	customerror "pokemon-game-api/pkgs/error"
)

type PokemonGateway struct{}

func NewPokemonGateway() IPokemonGateway {
	return PokemonGateway{}
}

func (p PokemonGateway) GetPokeapiPokemonDetailById(id string) (*PokemonGatewayResponse, error) {
	url := *config.PokeapiBaseUrl + "pokemon/" + id
	response, err := http.Get(url)

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokemonGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokemonGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	result := new(PokemonGatewayResponse)

	if err = json.Unmarshal(responseBody, result); err != nil {
		return nil, customerror.NewCustomError(constants.PokemonGwy,
			http.StatusInternalServerError, customerror.UnableToParseJsonToStruct)
	}

	return result, nil
}
