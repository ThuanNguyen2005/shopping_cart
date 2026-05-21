package v1dto

//data transform obj
type UserDTO struct {
	UUID   string `json:"uuid"`
	Name   string `json:"full_name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Status string `json:"status"`
	Level  string `json:"level"`
}

type UserInputCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_address"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_regex"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UserInputUpdateDTO struct {
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email,email_address"`
	Age      int    `json:"age" binding:"omitempty,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_regex"`
	Status   int    `json:"status" binding:"omitempty,oneof=1 2"`
	Level    int    `json:"level" binding:"omitempty,oneof=1 2"`
}

func MapUserDTO() *UserDTO {
	return &UserDTO{}
}

func MapUsersDTO() *[]UserDTO {
	return &[]UserDTO{}
}

func transformStatusToText(status int) string {
	switch status {
	case 1:
		return "show"
	case 2:
		return "hidden"
	default:
		return "none"
	}
}

func transformLevelToText(status int) string {
	switch status {
	case 1:
		return "admin"
	case 2:
		return "member"
	default:
		return "none"
	}
}

func (uc *UserInputCreateDTO) MapUserCreateToModel() {
}

func (uc *UserInputUpdateDTO) MapUserUpdateToModel() {
}
