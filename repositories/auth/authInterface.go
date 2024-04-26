package authrepo

import (
	"pokemon-game-api/domains/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthRepository interface {
	CreateUser(user *entities.User) (*mongo.InsertOneResult, error)
	ReadAllUser() (*[]entities.User, error)
	ReadUserById(id string) (*entities.User, error)
	UpdateUserById(id string, user *entities.User) (string, error)
	DeleteUserById(id string) (string, error)
}
