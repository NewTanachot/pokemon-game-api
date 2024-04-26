package pokedexctr

import (
	"pokemon-game-api/domains/models"
	pokedexusc "pokemon-game-api/usercases/pokedex"
)

// omitempty in json will hide json property in json response if this field have not data
type PokedexControllerResponse struct {
	Id          int                     `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Pokemons    []models.PokedexPokemon `json:"pokemons"`
}

func NewPokedexControllerResponse(usecase *pokedexusc.PokedexUsecaseResponse) *PokedexControllerResponse {
	result := PokedexControllerResponse{
		Id:          usecase.Id,
		Name:        usecase.Name,
		Description: usecase.Description,
	}

	for _, v := range usecase.Pokemons {
		result.Pokemons = append(result.Pokemons, models.PokedexPokemon{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &result
}
