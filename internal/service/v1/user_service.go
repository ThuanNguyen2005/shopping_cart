package v1service

import (
	"23.usersapi/internal/repository"
)

type serviceUser struct {
	repoUser repository.UserRepository
}

func NewServiceUser(repoUser repository.UserRepository) ServiceUser {
	return &serviceUser{
		repoUser: repoUser,
	}
}

func (us *serviceUser) GetAllUsers(search string, page, limit int) error {
	return nil
}

func (us *serviceUser) PostUser() error {
	return nil
}

func (us *serviceUser) GetUserByUUID(uuid string) error {
	return nil
}

func (us *serviceUser) PutUser(uuid string) error {
	return nil
}

func (us *serviceUser) DeleteUser(uuid string) error {
	return nil
}
