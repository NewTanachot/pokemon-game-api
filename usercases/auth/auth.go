package authusc

import (
	"pokemon-game-api/domains/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	UserName    string
	DisplayName string
	Password    string
}

type UserResponse struct {
	Id          primitive.ObjectID
	UserName    string
	DisplayName string
	Password    string
	IvKey       string
	Pokemons    []entities.Pokemon
	CreateAt    primitive.Timestamp
}
