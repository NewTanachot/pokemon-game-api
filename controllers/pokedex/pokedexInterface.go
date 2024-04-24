package pokedexctr

import "github.com/gin-gonic/gin"

type IPokedexController interface {
	GetPokedexDetail(c *gin.Context)
}
