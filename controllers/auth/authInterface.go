package authctr

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
}
