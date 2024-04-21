package pokemonusc

type IPokemonUsecase interface {
	GetPokemonDetailById(id string) (*PokemonUsecaseResponse, error)
}
