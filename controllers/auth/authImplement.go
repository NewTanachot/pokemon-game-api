package authctr

import (
	"net/http"
	"pokemon-game-api/domains/models"
	"pokemon-game-api/pkgs/constants"
	customerror "pokemon-game-api/pkgs/error"
	authusc "pokemon-game-api/usercases/auth"

	"github.com/gin-gonic/gin"
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
			http.StatusBadRequest, customerror.InvalidInput)

		c.AbortWithStatusJSON(cErr.Status, cErr.GetError())
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
	}

	c.JSON(http.StatusOK, result)
}

func (a AuthController) Login(c *gin.Context) {

}

func (a AuthController) GetAllUser(c *gin.Context) {
	uscResponse, cErr := a.AuthUsecase.GetAllUser()

	if cErr != nil {
		// pErr := customerror.ParseFrom(cErr)
		// c.AbortWithStatusJSON(pErr.Status, pErr.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, cErr.Error())
	}

	result := new([]UserResponse)
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

		*result = append(*result, UserResponse{
			Id:          v.Id,
			UserName:    v.UserName,
			DisplayName: v.DisplayName,
			Password:    v.Password,
			IvKey:       v.IvKey,
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
	}

	result := UserResponse{
		Id:          uscResp.Id,
		UserName:    uscResp.UserName,
		DisplayName: uscResp.DisplayName,
		Password:    uscResp.Password,
		IvKey:       uscResp.IvKey,
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
