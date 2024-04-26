package authusc

import "go.mongodb.org/mongo-driver/mongo"

type IAuthUsecase interface {
	CreateUser(user *CreateUserRequest) (*mongo.InsertOneResult, error)
	GetAllUser() (*[]UserResponse, error)
	GetUserById(id string) (*UserResponse, error)
}
