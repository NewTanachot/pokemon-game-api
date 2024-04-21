package pokedexusc

type IPokedexUsecase interface {
	GetPokedex(region string) (*PokedexUsecaseResponse, error)
}
