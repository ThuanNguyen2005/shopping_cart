package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ERR_CODE_BAD_REQUEST     ErrorCode = "BAD_REQUEST"
	ERR_CODE_NOT_FOUND       ErrorCode = "NOT_FOUND"
	ERR_CODE_CONFLICT        ErrorCode = "CONFLICT"
	ERR_CODE_INTERNAL_SERVER ErrorCode = "INTERNAL_SERVER"
)

type AppError struct {
	message string
	code    string
	Err     error
}

func NewAppError(message string, code string) *AppError {
	return &AppError{
		message: message,
		code:    code,
	}
}

func WrapAppError(message string, code string, err error) error {
	return &AppError{
		message: message,
		code:    code,
		Err:     err,
	}
}

func (ae *AppError) Error() string {
	return ""
}

func ResponseError(ctx *gin.Context, err error) {
	if appError, ok := err.(*AppError); ok {
		respond := gin.H{
			"error": appError.message,
			"code":  appError.code,
		}
		if appError.Err != nil {
			respond["detail"] = appError.Err.Error()
		}
		ctx.JSON(httpStatusFromCode(ErrorCode(appError.code)), respond)
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		"code":  ERR_CODE_INTERNAL_SERVER,
	})
}

func ResponseValidatorError(ctx *gin.Context, validator any) {
	ctx.JSON(http.StatusBadRequest, validator)
}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ERR_CODE_BAD_REQUEST:
		return http.StatusBadRequest
	case ERR_CODE_NOT_FOUND:
		return http.StatusNotFound
	case ERR_CODE_CONFLICT:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
