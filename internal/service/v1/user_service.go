package v1service

import (
	"database/sql"
	"errors"
	"strconv"

	"23.usersapi/internal/db/sqlc"
	"23.usersapi/internal/repository"
	"23.usersapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type serviceUser struct {
	repoUser repository.UserRepository
}

func NewServiceUser(repoUser repository.UserRepository) ServiceUser {
	return &serviceUser{
		repoUser: repoUser,
	}
}

func (us *serviceUser) GetAllUsers(context *gin.Context, search, sort, orderBy string, limit, page int32) ([]sqlc.User, int32, error) {
	ctx := context.Request.Context()
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limitENV, err := strconv.Atoi(utils.Getenv("LIMIT_IN_ONE_PAGE", "10"))
		if err != nil {
			limitENV = 10
		}
		limit = int32(limitENV)
	}
	if orderBy == "" {
		orderBy = "create_at"
	}
	if sort == "" {
		sort = "desc"
	}
	offset := (page - 1) * limit
	users, err := us.repoUser.FindAllUser(ctx, search, sort, orderBy, limit, page, offset)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapAppError("Cannot display all users", string(utils.ERR_CODE_INTERNAL_SERVER), err)
	}
	total, err := us.repoUser.UserTotal(ctx, search)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapAppError("Cannot find total user", string(utils.ERR_CODE_INTERNAL_SERVER), err)
	}
	return users, int32(total), nil
}

func (us *serviceUser) PostUser(ctx *gin.Context, userInput sqlc.CreateUserParams) (sqlc.User, error) {
	userInput.UserEmail = utils.NormalizeString(userInput.UserEmail)
	if us.repoUser.FindEmail(userInput.UserEmail) {
		return sqlc.User{}, utils.NewAppError("Email is exist", string(utils.ERR_CODE_CONFLICT))
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapAppError("Password cannot convert to hash password.", string(utils.ERR_CODE_INTERNAL_SERVER), err)
	}
	userInput.UserPassword = string(hashPassword)
	user, err := us.repoUser.CreateUser(ctx.Request.Context(), userInput)
	if err != nil {
		return sqlc.User{}, utils.WrapAppError("Cannot create user", string(utils.ERR_CODE_INTERNAL_SERVER), err)
	}
	return user, nil
}

func (us *serviceUser) GetUserByUUID(uuid string) error {
	return nil
}

func (us *serviceUser) PutUser(ctx *gin.Context, userInput sqlc.UpdateUserParams) (sqlc.User, error) {
	if userInput.UserPassword != nil && *userInput.UserPassword != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(*userInput.UserPassword), bcrypt.DefaultCost)
		if err != nil {
			return sqlc.User{}, utils.WrapAppError("Password cannot convert to hash password.", string(utils.ERR_CODE_INTERNAL_SERVER), err)
		}
		hashed := string(hashPassword)
		userInput.UserPassword = &hashed
	}
	user, err := us.repoUser.UpdateUser(ctx.Request.Context(), userInput)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewAppError("Not Found UUID", string(utils.ERR_CODE_NOT_FOUND))
		}
		return sqlc.User{}, utils.NewAppError("Failed to update user", string(utils.ERR_CODE_BAD_REQUEST))
	}
	return user, nil
}

func (us *serviceUser) SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := us.repoUser.SoftDeleteUser(ctx.Request.Context(), uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewAppError("UUID Not Found", string(utils.ERR_CODE_NOT_FOUND))
		}
		return sqlc.User{}, utils.NewAppError("Failed to delete user", string(utils.ERR_CODE_BAD_REQUEST))
	}
	return user, nil
}

func (us *serviceUser) RecoverUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := us.repoUser.RecoverUser(ctx.Request.Context(), uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewAppError("UUID Not Found", string(utils.ERR_CODE_NOT_FOUND))
		}
		return sqlc.User{}, utils.NewAppError("Failed to delete user", string(utils.ERR_CODE_BAD_REQUEST))
	}
	return user, nil
}

func (us *serviceUser) DeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := us.repoUser.DeleteUser(ctx.Request.Context(), uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewAppError("UUID Not Found", string(utils.ERR_CODE_NOT_FOUND))
		}
		return sqlc.User{}, utils.NewAppError("Failed to delete user", string(utils.ERR_CODE_BAD_REQUEST))
	}
	return user, nil
}
