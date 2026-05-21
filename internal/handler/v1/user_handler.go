package v1handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	v1dto "23.usersapi/internal/dto/v1"
	v1service "23.usersapi/internal/service/v1"
	"23.usersapi/internal/utils"
	"23.usersapi/internal/validation"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	serviceUser v1service.ServiceUser
}

type GetUserByUUID struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetUserParams struct {
	Search string `form:"search" binding:"omitempty,search,min=1,max=100"`
	Page   int    `form:"page" binding:"omitempty,gt=0,lt=100"`
	Limit  int    `form:"limit" binding:"omitempty,gt=0,lt=100"`
}

func NewHandlerUser(serviceUser v1service.ServiceUser) *HandlerUser {
	return &HandlerUser{
		serviceUser: serviceUser,
	}
}

func (uh *HandlerUser) GetAllUsers(ctx *gin.Context) {
	var params GetUserParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}
	log.Println("Start")
	time.Sleep(10 * time.Second)
	log.Println("End")
	utils.RespondSuccess(ctx, http.StatusOK, "")
}

func (uh *HandlerUser) PostUser(ctx *gin.Context) {
	var input v1dto.UserInputCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	utils.RespondSuccess(ctx, http.StatusCreated, "")
}

func (uh *HandlerUser) GetUserByUUID(ctx *gin.Context) {
	var param GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "")
}

func (uh *HandlerUser) PutUser(ctx *gin.Context) {
	var param GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}

	var input v1dto.UserInputUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	utils.RespondSuccess(ctx, http.StatusOK, "")
}

func (uh *HandlerUser) DeleteUser(ctx *gin.Context) {
	var param GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	utils.RespondStatus(ctx, http.StatusNoContent)
}
func (uh *HandlerUser) Panic(ctx *gin.Context) {
	var temp []string
	fmt.Println(temp[1])
}
