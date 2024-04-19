package pokedexusc

type IPokedexUsecase interface {
	GetPokedex() (*PokedexUsecaseResponse, error)
}
