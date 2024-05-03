package authusc

import (
	"pokemon-game-api/domains/entities"
	authrepo "pokemon-game-api/repositories/auth"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthUsecase struct {
	AuthRepository authrepo.IAuthRepository
}

func NewAuthUsecase(authRepo authrepo.IAuthRepository) IAuthUsecase {
	return AuthUsecase{AuthRepository: authRepo}
}

func (a AuthUsecase) CreateUser(user *CreateUserRequest) (*mongo.InsertOneResult, error) {

	// TODO

	userEnt := entities.User{
		UserName:    user.UserName,
		DisplayName: user.DisplayName,
		Password:    user.Password,
		// CreateAt:    primitive.Timestamp{T: uint32(time.Now().Unix())},
		BaseEntity: entities.BaseEntity{
			CreateAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	result, cErr := a.AuthRepository.CreateUser(&userEnt)

	if cErr != nil {
		return nil, cErr
	}

	return result, nil
}

func (a AuthUsecase) GetAllUser() (*[]UserResponse, error) {
	response, cErr := a.AuthRepository.ReadAllUser()

	if cErr != nil {
		return nil, cErr
	}

	result := new([]UserResponse)

	for _, v := range *response {
		*result = append(*result, UserResponse{
			Id:          v.Id,
			UserName:    v.UserName,
			DisplayName: v.DisplayName,
			Password:    v.Password,
			IvKey:       v.IvKey,
			Pokemons:    v.Pokemons,
			CreateAt:    v.CreateAt,
		})
	}

	return result, nil
}

func (a AuthUsecase) GetUserById(id string) (*UserResponse, error) {
	repoResp, cErr := a.AuthRepository.ReadUserById(id)

	if cErr != nil {
		return nil, cErr
	}

	result := UserResponse{
		Id:          repoResp.Id,
		UserName:    repoResp.UserName,
		DisplayName: repoResp.DisplayName,
		Password:    repoResp.Password,
		IvKey:       repoResp.IvKey,
		Pokemons:    repoResp.Pokemons,
	}

	return &result, nil
}
