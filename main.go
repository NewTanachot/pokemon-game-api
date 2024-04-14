package main

import "pokemon-game-api/pkgs/server"

func main() {
	server.GinSetup()
	server.GinStart()
}
