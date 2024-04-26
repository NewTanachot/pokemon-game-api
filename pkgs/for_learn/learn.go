package forlearn

import (
	"pokemon-game-api/pkgs/di"

	"github.com/gin-gonic/gin"
)

var (
	learnCollection = di.MongoDb.Client.Database("learns").Collection("pokemon")
)

func CreatePokemon(c *gin.Context) {

}
