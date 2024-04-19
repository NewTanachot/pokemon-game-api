package pokedexgwy

type IPokedexGateway interface {
	GetPokeapiPokedex() (*PokedexGatewayResponse, error)
}
