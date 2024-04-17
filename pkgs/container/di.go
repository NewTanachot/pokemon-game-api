package container

import (
	pokedexctr "pokemon-game-api/controllers/pokedex"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	customlog "pokemon-game-api/pkgs/logs"
	pokedexusc "pokemon-game-api/usercases/pokedex"
	"sync"
)

var (
	once sync.Once

	// pokedex
	PokedexController pokedexctr.IPokedexController
	PokedexUsecase    pokedexusc.IPokedexUsecase
	PokedexGateway    pokedexgwy.IPokedexGateway
)

func AddDependencyInjections() {
	once.Do(func() {
		// pokedex
		PokedexGateway = pokedexgwy.NewPokedexGateway()
		PokedexUsecase = pokedexusc.NewPokedexUsecase(PokedexGateway)
		PokedexController = pokedexctr.NewPokedexController(PokedexUsecase)

		customlog.WriteBorderedInfoLog("Add Dependencies to DI container")
	})
}
