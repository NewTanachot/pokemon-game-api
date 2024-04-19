package customerror

import (
	"fmt"
	customlog "pokemon-game-api/pkgs/logs"
)

type CustomError struct {
	Package      string `json:"package"`
	ErrorCode    int    `json:"errorCode"`
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
}

func NewCustomError(pkg string, errCode int, msg string) error {
	errMsg := fmt.Sprintf("error occur in %v. - %v.", pkg, msg)
	customlog.WriteBorderedErrorLog(errMsg)

	return CustomError{
		Package:      pkg,
		ErrorCode:    errCode,
		Message:      msg,
		ErrorMessage: errMsg,
	}
}

func (e CustomError) Error() string {
	return e.ErrorMessage
}
