package v1service

type ServiceUser interface {
	GetAllUsers(search string, page, limit int) error
	PostUser() error
	GetUserByUUID(uuid string) error
	PutUser(uuid string) error
	DeleteUser(uuid string) error
}
