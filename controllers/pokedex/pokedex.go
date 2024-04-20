package pokedexctr

import (
	"pokemon-game-api/domains/models"
	pokedexusc "pokemon-game-api/usercases/pokedex"
)

type PokedexControllerResponse struct {
	Id          int                           `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Pokemons    []models.PokedexPokemonDetail `json:"pokemons"`
}

func NewPokedexControllerResponse(usecase *pokedexusc.PokedexUsecaseResponse) *PokedexControllerResponse {
	result := PokedexControllerResponse{
		Id:          usecase.Id,
		Name:        usecase.Name,
		Description: usecase.Description,
	}

	for _, v := range usecase.Pokemons {
		result.Pokemons = append(result.Pokemons, models.PokedexPokemonDetail{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &result
}
