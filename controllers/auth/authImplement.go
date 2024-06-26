package authctr

import (
	"net/http"
	"pokemon-game-api/domains/models"
	"pokemon-game-api/pkgs/constants"
	customerror "pokemon-game-api/pkgs/error"
	stringutils "pokemon-game-api/pkgs/utils/string"
	authusc "pokemon-game-api/usercases/auth"
	"time"

	"pokemon-game-api/pkgs/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	AuthUsecase authusc.IAuthUsecase
}

func NewAuthController(authUsc authusc.IAuthUsecase) IAuthController {
	return AuthController{AuthUsecase: authUsc}
}

func (a AuthController) Register(c *gin.Context) {
	ctrRequest := new(RegisterRequest)
	if err := c.BindJSON(ctrRequest); err != nil {
		cErr := customerror.NewCustomError(constants.AuthColl,
			http.StatusBadRequest, customerror.InvalidInput+err.Error())

		c.AbortWithStatusJSON(cErr.Status, cErr.GetError())
		return
	}

	uscRequest := authusc.CreateUserRequest{
		UserName:    ctrRequest.UserName,
		DisplayName: ctrRequest.DisplayName,
		Password:    ctrRequest.Password,
	}

	result, cErr := a.AuthUsecase.CreateUser(&uscRequest)

	if cErr != nil {
		pErr := customerror.ParseFrom(cErr)
		c.AbortWithStatusJSON(pErr.Status, pErr.GetError())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (a AuthController) Login(c *gin.Context) {
	req := new(LoginRequest)
	if err := c.BindJSON(req); err != nil {
		cErr := customerror.NewCustomError(constants.AuthColl,
			http.StatusBadRequest, customerror.InvalidInput+err.Error())

		c.AbortWithStatusJSON(cErr.Status, cErr.GetError())
		return
	}

	uscResp, cErr := a.AuthUsecase.GetUserById(req.Id)

	if cErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, cErr.Error())
		return
	}

	// TODO: move to usecase
	if uscResp.Password != req.Password {
		cErr := customerror.NewCustomError(constants.AuthColl,
			http.StatusBadRequest, customerror.WrongPassword)

		c.AbortWithStatusJSON(cErr.Status, cErr.GetError())
		return
	}

	// TODO: move to usecase
	jwt, err := createJwtToken(uscResp.Id)

	if err != nil {
		cErr := customerror.NewCustomError(constants.AuthColl,
			http.StatusBadRequest, customerror.UnableToCreateJWT+" "+err.Error())

		c.AbortWithStatusJSON(cErr.Status, cErr.GetError())
		return
	}

	c.JSONP(http.StatusOK, jwt)
}

func (a AuthController) GetAllUser(c *gin.Context) {
	uscResponse, cErr := a.AuthUsecase.GetAllUser()

	if cErr != nil {
		// pErr := customerror.ParseFrom(cErr)
		// c.AbortWithStatusJSON(pErr.Status, pErr.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, cErr.Error())
		return
	}

	result := []UserResponse{}

	// for _, v := range *uscResponse {
	// 	result = append(result, UserResponse{
	// 		Id: v.Id,
	// 		UserName: v.UserName,
	// 		DisplayName: v.DisplayName,
	// 		Password: v.Password,
	// 		IvKey: v.IvKey,
	// 		Pokemons: func() []PokemonDto {
	// 			pokeTemp := []PokemonDto{}
	// 			for _, j := range v.Pokemons {
	// 				pokeTemp = append(pokeTemp, PokemonDto{
	// 					Id: j.Id,
	// 					Name: j.Name,
	// 					Nickname: j.Nickname,
	// 					Level: j.Level,
	// 					Sequence: j.Sequence,
	// 					Moves: func() []models.Move {
	// 						moveTemp := []models.Move{}
	// 						for _, k := range j.Moves {
	// 							moveTemp = append(moveTemp, models.Move{
	// 								Id: k.Id,
	// 								Sequence: k.Sequence,
	// 								Name: k.Name,
	// 								Type: k.Type,
	// 							})
	// 						}
	// 						return moveTemp
	// 					}(),
	// 				})
	// 			}
	// 			return pokeTemp
	// 		}(),
	// 	})
	// }

	for _, v := range *uscResponse {
		PokemonDtos := []PokemonDto{}

		for _, j := range v.Pokemons {
			moveDtos := []models.Move{}

			for _, k := range j.Moves {
				moveDtos = append(moveDtos, models.Move{
					Id:       k.Id,
					Sequence: k.Sequence,
					Name:     k.Name,
					Type:     k.Type,
				})
			}

			PokemonDtos = append(PokemonDtos, PokemonDto{
				Id:       j.Id,
				Name:     j.Name,
				Nickname: j.Nickname,
				Level:    j.Level,
				Sequence: j.Sequence,
				Moves:    moveDtos,
			})
		}

		result = append(result, UserResponse{
			Id:          v.Id,
			UserName:    v.UserName,
			DisplayName: v.DisplayName,
			Password:    v.Password,
			IvKey:       v.IvKey,
			CreateAt:    v.CreateAt.T,
			Pokemons:    PokemonDtos,
		})
	}

	c.JSON(http.StatusOK, result)
}

func (a AuthController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	uscResp, cErr := a.AuthUsecase.GetUserById(id)

	if cErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, cErr.Error())
		return
	}

	result := UserResponse{
		Id:          uscResp.Id,
		UserName:    uscResp.UserName,
		DisplayName: uscResp.DisplayName,
		Password:    uscResp.Password,
		IvKey:       uscResp.IvKey,
		CreateAt:    uscResp.CreateAt.T,
		Pokemons: func() []PokemonDto {
			tempPokemon := []PokemonDto{}
			for _, v := range uscResp.Pokemons {
				tempPokemon = append(tempPokemon, PokemonDto{
					Id:       v.Id,
					Name:     v.Name,
					Nickname: v.Nickname,
					Level:    v.Level,
					Sequence: v.Sequence,
					Moves: func() []models.Move {
						tempMove := []models.Move{}
						for _, m := range v.Moves {
							tempMove = append(tempMove, models.Move{
								Id:       m.Id,
								Sequence: m.Sequence,
								Name:     m.Name,
								Type:     m.Type,
							})
						}

						return tempMove
					}(),
				})
			}

			return tempPokemon
		}(),
	}

	c.JSON(http.StatusOK, result)
}

// TODO: move to usecase or pkg
func createJwtToken(userId primitive.ObjectID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["user_id"] = userId
	claim["exp"] = time.Now().Add(168 * time.Hour).Unix()

	secretKeyBytes := []byte(*config.SecretKey)
	resultJwt, err := token.SignedString(secretKeyBytes)

	if err != nil {
		return stringutils.Empty, err
	}

	return resultJwt, nil
}
