package repository

import (
	"fmt"
)

type SqlRepositoryUser struct {
}

func NewInMemoryRepository() UserRepository {
	return &SqlRepositoryUser{}
}

func (ur *SqlRepositoryUser) FindAllUser() error {
	return nil
}
func (ur *SqlRepositoryUser) CreateUser() error {
	return nil
}
func (ur *SqlRepositoryUser) FindUserByUUID(uuid string) bool {
	return false
}
func (ur *SqlRepositoryUser) UpdateUser(uuid string) error {
	return fmt.Errorf("Failed to update user.")
}
func (ur *SqlRepositoryUser) DeleteUser(uuid string) bool {
	return false
}

func (ur *SqlRepositoryUser) FindEmail(email string) bool {
	return false
}
