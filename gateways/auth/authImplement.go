package authgwy

type IAuthGateway interface {
	CreateUser()
	ReadUserById()
	UpdateUserById()
	DeleteUserById()
}
