package v1service

import (
	"23.usersapi/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceUser interface {
	GetAllUsers(ctx *gin.Context, search, sort, orderBy string, limit, page int32) ([]sqlc.User, int32, error)
	PostUser(ctx *gin.Context, userInput sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByUUID(uuid string) error
	PutUser(ctx *gin.Context, userInput sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
	SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
	RecoverUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
}
