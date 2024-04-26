package authusc

import "pokemon-game-api/domains/entities"

type CreateUserRequest struct {
	UserName    string
	DisplayName string
	Password    string
}

type UserResponse struct {
	Id          string
	UserName    string
	DisplayName string
	Password    string
	IvKey       string
	Pokemons    []entities.Pokemon
}
