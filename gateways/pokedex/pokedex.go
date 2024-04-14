package pokedexgwy

type IPokedexGateway interface {
	GetPokeapiSinnohPokedex() (any, error)
}

type PokedexGateway struct{}
