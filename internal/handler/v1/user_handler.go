package v1handler

import (
	"net/http"

	v1dto "23.usersapi/internal/dto/v1"
	v1service "23.usersapi/internal/service/v1"
	"23.usersapi/internal/utils"
	"23.usersapi/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerUser struct {
	serviceUser v1service.ServiceUser
}

func NewHandlerUser(serviceUser v1service.ServiceUser) *HandlerUser {
	return &HandlerUser{
		serviceUser: serviceUser,
	}
}

func (uh *HandlerUser) GetAllUsers(ctx *gin.Context) {
	var params v1dto.GetUserParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}

	users, total, err := uh.serviceUser.GetAllUsers(ctx, params.Search, params.Sort, params.OrderBy, params.Limit, params.Page)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := v1dto.MapModelToUsersDTO(users)
	userPaginInfo := v1dto.DataPaginInfoCombine(userDTO, params.Page, params.Limit, total)
	utils.RespondSuccess(ctx, http.StatusOK, "export all user successful", userPaginInfo)
}

func (uh *HandlerUser) PostUser(ctx *gin.Context) {
	var input v1dto.UserInputCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	userInput := input.MapUserCreateToModel()
	user, err := uh.serviceUser.PostUser(ctx, userInput)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	createUser := v1dto.MapModelToUserDTO(user)
	utils.RespondSuccess(ctx, http.StatusCreated, "Created user successfully", createUser)
}

func (uh *HandlerUser) GetUserByUUID(ctx *gin.Context) {
	var param v1dto.GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "")
}

func (uh *HandlerUser) PutUser(ctx *gin.Context) {
	var paramUUID v1dto.GetUserByUUID
	if err := ctx.ShouldBindUri(&paramUUID); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	userUUID, err := uuid.Parse(paramUUID.Uuid)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	var input v1dto.UserInputUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	userInput := input.MapUserUpdateToModel(userUUID)
	user, err := uh.serviceUser.PutUser(ctx, userInput)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	userOutput := v1dto.MapModelToUserDTO(user)
	utils.RespondSuccess(ctx, http.StatusOK, "Updated user successfully", userOutput)
}

func (uh *HandlerUser) DeleteUser(ctx *gin.Context) {
	var param v1dto.GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	paramUUID, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	user, err := uh.serviceUser.DeleteUser(ctx, paramUUID)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := v1dto.MapModelToUserDTO(user)
	utils.RespondSuccess(ctx, http.StatusOK, "Deleted user successfully", userDTO)
}

func (uh *HandlerUser) RecoverUser(ctx *gin.Context) {
	var param v1dto.GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	paramUUID, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	user, err := uh.serviceUser.RecoverUser(ctx, paramUUID)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := v1dto.MapModelToUserDTO(user)
	utils.RespondSuccess(ctx, http.StatusOK, "Recover user successfully", userDTO)
}

func (uh *HandlerUser) SoftDeleteUser(ctx *gin.Context) {
	var param v1dto.GetUserByUUID
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	paramUUID, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseValidatorError(ctx, validation.HandlerValidation(err))
		return
	}
	user, err := uh.serviceUser.SoftDeleteUser(ctx, paramUUID)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := v1dto.MapModelToUserDTO(user)
	utils.RespondSuccess(ctx, http.StatusOK, "Deleted user successfully", userDTO)
}

// func (uh *HandlerUser) Panic(ctx *gin.Context) {
// 	var temp []string
// 	fmt.Println(temp[1])
// }
