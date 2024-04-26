package authusc

import (
	"pokemon-game-api/domains/entities"
	authrepo "pokemon-game-api/repositories/auth"

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
