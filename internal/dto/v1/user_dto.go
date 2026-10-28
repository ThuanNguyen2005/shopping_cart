package v1dto

import (
	"strconv"

	"23.usersapi/internal/db/sqlc"
	"23.usersapi/internal/utils"
	"github.com/google/uuid"
)

// data transform obj
type UserDTO struct {
	ID       int       `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"full_name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Age      *int      `json:"age"`
	Status   string    `json:"status"`
	Level    string    `json:"level"`
	CreateAt string    `json:"create_at"`
}

type UserInputCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_address"`
	Age      int32  `json:"age" binding:"omitempty,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_regex"`
	Status   int    `json:"status" binding:"required,oneof=1 2 3"`
	Level    int    `json:"level" binding:"required,oneof=1 2 3"`
}

type UserInputUpdateDTO struct {
	Name     *string `json:"name" binding:"omitempty"`
	Age      *int32  `json:"age" binding:"omitempty,gt=0"`
	Password *string `json:"password" binding:"omitempty,min=8,password_regex"`
	Status   *int32  `json:"status" binding:"omitempty,oneof=1 2 3"`
	Level    *int32  `json:"level" binding:"omitempty,oneof=1 2 3"`
}

func MapUserDTO() *UserDTO {
	return &UserDTO{}
}

func MapModelToUsersDTO(users []sqlc.User) *[]UserDTO {
	var userDTO []UserDTO
	for _, user := range users {
		userDTO = append(userDTO, MapModelToUserDTO(user))
	}
	return &userDTO
}

func transformStatusToText(status int) string {
	switch status {
	case 1:
		return "ACTIVE"
	case 2:
		return "INACTIVE"
	case 3:
		return "BANNED"
	default:
		return "NONE"
	}
}

func transformLevelToText(status int) string {
	switch status {
	case 1:
		return "ADMIN"
	case 2:
		return "VIP"
	case 3:
		return "MEMBER"
	default:
		return "NONE"
	}
}

func (uc *UserInputCreateDTO) MapUserCreateToModel() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		UserName:     uc.Name,
		UserEmail:    uc.Email,
		UserPassword: uc.Password,
		UserAge:      utils.ConvertToInt32Pointer(uc.Age),
		UserStatus:   int32(uc.Status),
		UserLevel:    int32(uc.Level),
	}
}

func MapModelToUserDTO(user sqlc.User) UserDTO {
	dto := UserDTO{
		ID:       int(user.UserID),
		UUID:     user.UserUuid,
		Name:     user.UserName,
		Email:    user.UserEmail,
		Password: user.UserPassword,
		Status:   transformStatusToText(int(user.UserStatus)),
		Level:    transformLevelToText(int(user.UserLevel)),
		CreateAt: user.CreateAt.Format("02-01-2006 15:04:05"),
	}
	if user.UserAge != nil {
		age := int(*user.UserAge)
		dto.Age = &age
	}
	return dto
}

func (uc *UserInputUpdateDTO) MapUserUpdateToModel(userUUID uuid.UUID) sqlc.UpdateUserParams {
	return sqlc.UpdateUserParams{
		UserName:     uc.Name,
		UserPassword: uc.Password,
		UserAge:      uc.Age,
		UserStatus:   uc.Status,
		UserLevel:    uc.Level,
		UserUuid:     userUUID,
	}
}

type GetUserByUUID struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetUserParams struct {
	Search  string `form:"search" binding:"omitempty,search,min=1,max=100"`
	Page    int32  `form:"page" binding:"omitempty,gt=0"`
	Limit   int32  `form:"limit" binding:"omitempty,gt=0,lt=500"`
	Offset  int32  `form:"offset" binding:"omitempty,gt=0"`
	Sort    string `form:"sort" binding:"omitempty,oneof=desc asc"`
	OrderBy string `form:"order_by" binding:"omitempty,oneof=create_at user_id"`
}

type PaginInfo struct {
	Data  any
	Pagin Pagination
}

type Pagination struct {
	Page        int32
	Limit       int32
	RecordTotal int32
	PageTotal   int32
	HasNext     bool
	HasPre      bool
}

func DataPaginInfoCombine(data any, page int32, limit int32, recordTotal int32) PaginInfo {
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
	pageTotal := (recordTotal + limit - 1) / limit
	return PaginInfo{
		Data: data,
		Pagin: Pagination{
			Page:        page,
			Limit:       limit,
			HasPre:      page > 1,
			HasNext:     page < pageTotal,
			RecordTotal: recordTotal,
			PageTotal:   pageTotal,
		},
	}
}
