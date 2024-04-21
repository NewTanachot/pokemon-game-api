package pokedexgwy

import (
	"encoding/json"
	"io"
	"net/http"
	"pokemon-game-api/pkgs/config"
	"pokemon-game-api/pkgs/constants"
	customerror "pokemon-game-api/pkgs/error"
)

type PokedexGateway struct{}

func NewPokedexGateway() IPokedexGateway {
	return PokedexGateway{}
}

func (p PokedexGateway) GetPokeapiPokedex(regionNo string) (*PokedexGatewayResponse, error) {
	url := config.PokeapiBaseUrl + "pokedex/" + regionNo
	response, err := http.Get(url)

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.InvalidResponse)
	}

	result := new(PokedexGatewayResponse)

	if err = json.Unmarshal(responseBody, result); err != nil {
		return nil, customerror.NewCustomError(constants.PokedexGwy,
			http.StatusInternalServerError, customerror.UnableToParseJsonToStruct)
	}

	return result, nil
}
