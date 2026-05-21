package repository

type UserRepository interface {
	FindAllUser() error
	CreateUser() error
	FindUserByUUID(uuid string) bool
	UpdateUser(uuid string) error
	DeleteUser(uuid string) bool
	FindEmail(email string) bool
}
