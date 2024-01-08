package userController

import (
	"github.com/yasar-arafat/go-project-starter/dal"
	"github.com/yasar-arafat/go-project-starter/model"
)

// TODO: Add logging and Error handling.
func CreateUser(u *model.User) error {
	userStore := dal.UserStore{}
	if err := userStore.Create(u); err != nil {
		return err
	}
	return nil
}

// TODO: add logging and error handling
func GetAllUsers() ([]model.User, error) {
	userStore := dal.UserStore{}
	users, err := userStore.GetAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
