package pokemongwy

type IPokemonGateway interface {
	GetPokeapiPokemonDetailById(id string) (*PokemonGatewayResponse, error)
}
