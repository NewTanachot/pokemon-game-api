package dependency

import (
	pokedexctr "pokemon-game-api/controllers/pokedex"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	pokedexusc "pokemon-game-api/usercases/pokedex"
	"sync"
)

var (
	once       sync.Once
	isInjected bool = false

	// pokedex
	PokedexController pokedexctr.IPokedexController
	PokedexUsecase    pokedexusc.IPokedexUsecase
	PokedexGateway    pokedexgwy.IPokedexGateway
)

func UseDependencyInjection() {
	once.Do(func() {
		// pokedex
		PokedexGateway = pokedexgwy.NewPokedexGateway()
		PokedexUsecase = pokedexusc.NewPokedexUsecase(PokedexGateway)
		PokedexController = pokedexctr.NewPokedexController(PokedexUsecase)

		isInjected = true
	})
}

func IsInjected() bool {
	return isInjected
}
