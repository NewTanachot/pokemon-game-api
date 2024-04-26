package main

import (
	"pokemon-game-api/pkgs/di"
	"pokemon-game-api/pkgs/server"
)

func main() {
	defer di.MongoDb.CloseMongoDb()

	server.GinSetup()
	server.GinStart()
}
