package di

import (
	authctr "pokemon-game-api/controllers/auth"
	pokedexctr "pokemon-game-api/controllers/pokedex"
	pokemonctr "pokemon-game-api/controllers/pokemon"
	pokedexgwy "pokemon-game-api/gateways/pokedex"
	pokemongwy "pokemon-game-api/gateways/pokemon"
	"pokemon-game-api/pkgs/database"
	customlog "pokemon-game-api/pkgs/logs"
	authrepo "pokemon-game-api/repositories/auth"
	authusc "pokemon-game-api/usercases/auth"
	pokedexusc "pokemon-game-api/usercases/pokedex"
	pokemonusc "pokemon-game-api/usercases/pokemon"

	"sync"
)

// TODO - restructure to singleton like database package

var (
	once    sync.Once
	MongoDb *database.MongoDb
	// need to allocate pointer when init for access address of pointer after that.
	// ref https://edwinsiby.medium.com/runtime-error-invalid-memory-address-or-nil-pointer-dereference-golang-dd4a58ab7536
	PokedexController *pokedexctr.IPokedexController = new(pokedexctr.IPokedexController)
	PokemonController *pokemonctr.IPokemonController = new(pokemonctr.IPokemonController)
	AuthController    *authctr.IAuthController       = new(authctr.IAuthController)
)

func AddDependencyInjections() {
	once.Do(func() {
		// mongodb
		MongoDb = database.NewMongoDbClient()

		// auth
		authRepository := authrepo.NewAuthGatway(MongoDb.DbClient)
		authUsecase := authusc.NewAuthUsecase(authRepository)
		*AuthController = authctr.NewAuthController(authUsecase)

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
