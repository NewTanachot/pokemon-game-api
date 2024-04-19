package pokedexusc

import "pokemon-game-api/domains/models"

type PokedexUsecaseResponse struct {
	Id          uint
	Name        string
	Description string
	Pokemons    []models.PokedexPokemonDetail
}
