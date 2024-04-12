package main

import "pokemon-game-api/pkg/server"

func main() {
	server.GinSetup()
	server.GinStart()
}
