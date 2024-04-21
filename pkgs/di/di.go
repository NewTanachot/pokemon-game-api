package di

import (
	pokedexctr "pokemon-game-api/controllers/pokedex"
	pokemonctr "pokemon-game-api/controllers/pokemon"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	pokemongwy "pokemon-game-api/gateways/pokemon"
	customlog "pokemon-game-api/pkgs/logs"
	pokedexusc "pokemon-game-api/usercases/pokedex"
	pokemonusc "pokemon-game-api/usercases/pokemon"

	"sync"
)

var (
	// pokedex
	PokedexGateway    pokedexgwy.IPokedexGateway
	PokedexUsecase    pokedexusc.IPokedexUsecase
	PokedexController pokedexctr.IPokedexController

	// pokemon
	PokemonGateway    pokemongwy.IPokemonGateway
	PokemonUsecase    pokemonusc.IPokemonUsecase
	PokemonController pokemonctr.IPokemonController

	once sync.Once
)

func AddDependencyInjections() {
	once.Do(func() {
		// pokedex
		PokedexGateway = pokedexgwy.NewPokedexGateway()
		PokedexUsecase = pokedexusc.NewPokedexUsecase(PokedexGateway)
		PokedexController = pokedexctr.NewPokedexController(PokedexUsecase)

		// pokemon
		PokemonGateway = pokemongwy.NewPokemonGateway()
		PokemonUsecase = pokemonusc.NewPokemonUsecase(PokemonGateway)
		PokemonController = pokemonctr.NewPokemonControllor(PokemonUsecase)

		customlog.WriteBorderedInfoLog("Add Dependencies to DI container")
	})
}
