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
	once              sync.Once
	PokedexController *pokedexctr.IPokedexController
	PokemonController *pokemonctr.IPokemonController
)

func AddDependencyInjections() {
	once.Do(func() {
		// mongodb
		// mongoDb := database.NewMongoDbClient()

		// pokedex
		pokedexGateway := pokedexgwy.NewPokedexGateway()
		pokedexUsecase := pokedexusc.NewPokedexUsecase(pokedexGateway)
		*PokedexController = pokedexctr.NewPokedexController(pokedexUsecase)

		// pokemon
		pokemonGateway := pokemongwy.NewPokemonGateway()
		pokemonUsecase := pokemonusc.NewPokemonUsecase(pokemonGateway)
		*PokemonController = pokemonctr.NewPokemonControllor(pokemonUsecase)

		customlog.WriteBorderedInfoLog("Add Dependencies to DI container")
	})
}
