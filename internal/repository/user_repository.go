package repository

import (
	"context"
	"fmt"
	"log"

	"23.usersapi/internal/db"
	"23.usersapi/internal/db/sqlc"
	"github.com/google/uuid"
)

type SqlRepositoryUser struct {
	DB sqlc.Querier
}

func NewInMemoryRepository(DB sqlc.Querier) UserRepository {
	return &SqlRepositoryUser{
		DB: DB,
	}
}

func (ur *SqlRepositoryUser) FindAllUser(ctx context.Context, search, sort, orderBy string, limit, page, offset int32) ([]sqlc.User, error) {
	var (
		users []sqlc.User
		err   error
	)
	switch {
	case orderBy == "user_id" && sort == "desc":
		users, err = ur.DB.ListUsersIDDESC(ctx, sqlc.ListUsersIDDESCParams{
			Limit:  limit,
			Offset: offset,
			Search: &search,
		})
	case orderBy == "user_id" && sort == "asc":
		users, err = ur.DB.ListUsersIDASC(ctx, sqlc.ListUsersIDASCParams{
			Limit:  limit,
			Offset: offset,
			Search: &search,
		})
	case orderBy == "create_at" && sort == "desc":
		users, err = ur.DB.ListUsersCreateAtDESC(ctx, sqlc.ListUsersCreateAtDESCParams{
			Limit:  limit,
			Offset: offset,
			Search: &search,
		})
	case orderBy == "create_at" && sort == "asc":
		users, err = ur.DB.ListUsersCreateAtASC(ctx, sqlc.ListUsersCreateAtASCParams{
			Limit:  limit,
			Offset: offset,
			Search: &search,
		})
	}
	if err != nil {
		return []sqlc.User{}, err
	}
	log.Println(users)

	return users, nil
}

func (ur *SqlRepositoryUser) FindAllUserv2(ctx context.Context, search, sort, orderBy string, limit, page, offset int32) ([]sqlc.User, error) {
	query := `SELECT *
		FROM users
		WHERE delete_at IS NULL
		AND(
			$1::TEXT IS NULL
			OR $1::TEXT=''
			OR user_name ILIKE '%'|| $1 ||'%'
			OR user_email ILIKE '%'|| $1 ||'%'
		)`
	query += fmt.Sprintf(" ORDER By %s %s LIMIT %s OFFSET %s;", orderBy, sort, limit, offset)
	rows, err := db.DBpool.Query(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []sqlc.User{}
	for rows.Next() {
		var i sqlc.User
		if err := rows.Scan(
			&i.UserID,
			&i.UserUuid,
			&i.UserName,
			&i.UserEmail,
			&i.UserPassword,
			&i.UserAge,
			&i.UserStatus,
			&i.UserLevel,
			&i.CreateAt,
			&i.UpdateAt,
			&i.DeleteAt,
		); err != nil {
			return nil, err
		}
		users = append(users, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *SqlRepositoryUser) UserTotal(ctx context.Context, search string) (int64, error) {
	total, err := ur.DB.UserTotal(ctx, &search)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ur *SqlRepositoryUser) CreateUser(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := ur.DB.CreateUser(ctx, userParam)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}
func (ur *SqlRepositoryUser) FindUserByUUID(uuid string) bool {
	return false
}
func (ur *SqlRepositoryUser) UpdateUser(ctx context.Context, userParam sqlc.UpdateUserParams) (sqlc.User, error) {
	user, err := ur.DB.UpdateUser(ctx, userParam)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

func (ur *SqlRepositoryUser) RecoverUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := ur.DB.RecoveUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

func (ur *SqlRepositoryUser) SoftDeleteUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := ur.DB.SoftDeleteUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

func (ur *SqlRepositoryUser) DeleteUser(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := ur.DB.DeleteUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

func (ur *SqlRepositoryUser) FindEmail(email string) bool {
	return false
}
