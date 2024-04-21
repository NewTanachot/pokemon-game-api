package pokedexgwy

type IPokedexGateway interface {
	GetPokeapiPokedex(regionNo string) (*PokedexGatewayResponse, error)
}
