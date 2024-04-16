package main

import (
	"pokemon-game-api/pkgs/server"

	"github.com/joho/godotenv"
)

func main() {
	// load configuration from .env file
	godotenv.Load()

	server.GinSetup()
	server.GinStart()
}
