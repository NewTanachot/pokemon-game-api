package pokemon

type IPokemonGateway interface {
	GetPokeapiPokemonDetailById(id int)
}
