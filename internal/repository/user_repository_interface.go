package repository

import (
	"context"

	"23.usersapi/internal/db/sqlc"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindAllUserv2(ctx context.Context, search, sort, orderBy string, limit, page, offset int32) ([]sqlc.User, error)
	FindAllUser(ctx context.Context, search, sort, orderBy string, limit, page, offset int32) ([]sqlc.User, error)
	UserTotal(ctx context.Context, search string) (int64, error)
	CreateUser(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error)
	FindUserByUUID(uuid string) bool
	UpdateUser(ctx context.Context, userParam sqlc.UpdateUserParams) (sqlc.User, error)
	RecoverUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
	SoftDeleteUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
	DeleteUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
	FindEmail(email string) bool
}
