package customerror

import (
	"fmt"
	customlog "pokemon-game-api/pkgs/logs"
)

type CustomError struct {
	Status  int    `json:"status"`
	Module  string `json:"module"`
	Message string `json:"message"`
	Log     string `json:"log"`
}

func NewCustomError(module string, status int, msg string) CustomError {
	return CustomError{
		Status:  status,
		Module:  module,
		Message: msg,
		Log:     fmt.Sprintf("error occur in %v. - %v.", module, msg),
	}
}

func ParseFrom(err error) CustomError {
	return err.(CustomError)
}

func (e CustomError) GetError() error {
	e.WriteLog()
	return e
}

func (e CustomError) WriteLog() string {
	customlog.WriteBorderedErrorLog(e.Log)
	return e.Log
}

func (e CustomError) Error() string {
	e.WriteLog()
	return e.Log
}
