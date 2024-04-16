package dependency

import (
	pokedexctr "pokemon-game-api/controllers/pokedex"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	"pokemon-game-api/pkgs/logs"
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

func AddDependencyInjections() error {
	// todo use cahnnel to return error
	once.Do(func() {
		// pokedex
		PokedexGateway = pokedexgwy.NewPokedexGateway()
		PokedexUsecase = pokedexusc.NewPokedexUsecase(PokedexGateway)
		PokedexController = pokedexctr.NewPokedexController(PokedexUsecase)

		logs.WriteBorderedInfoLog("Add Dependencies to DI container")
	})

	return nil
}
