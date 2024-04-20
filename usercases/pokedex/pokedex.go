package pokedexusc

import "pokemon-game-api/domains/models"

type PokedexUsecaseResponse struct {
	Id          int
	Name        string
	Description string
	Pokemons    []models.PokedexPokemonDetail
}
