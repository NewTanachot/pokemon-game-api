package pokedexgwy

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewPokedexGateway() IPokedexGateway {
	return PokedexGateway{}
}

func (p PokedexGateway) GetPokeapiSinnohPokedex() (any, error) {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/5/")

	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	// result := new(repo_response.CovidCaseResponse)
	result := new(any)

	if err = json.Unmarshal(responseBody, result); err != nil {
		return nil, err
	}

	return result, nil
}
