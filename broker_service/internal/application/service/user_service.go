package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (us *UserService) CreateUser(newUser *model.UserDTO) error {
	if newUser == nil {
		return fmt.Errorf("user can't be null")
	}

	if err := rpc_client.Call(addrs.UserAddr, "UserService.CreateUser", &newUser, &struct{}{}); err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(user *model.UserDTO) error {
	if user == nil {
		return fmt.Errorf("user can't be null")
	}

	if err := rpc_client.Call(addrs.UserAddr, "UserService.UpdateUser", &user, &struct{}{}); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetAllUsers() ([]*model.UserDTO, error) {
	var users []*model.UserDTO

	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetAllUsers", &users, &struct{}{}); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetAllActiveUsers() ([]*model.UserDTO, error) {
	var users []*model.UserDTO

	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetAllActiveUsers", &users, &struct{}{}); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetUserById(userId int64) (*model.UserDTO, error) {
	var user *model.UserDTO

	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetUserById", &user, &userId); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByUsername(username string) (*model.UserDTO, error) {
	var user *model.UserDTO

	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetUserByUsername", &username, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) DeactivateUser(userId int64) error {
	if err := rpc_client.Call(addrs.UserAddr, "UserService.DeactivateUser", &userId, &struct{}{}); err != nil {
		return err
	}

	return nil
}

func (us *UserService) ReactivateUser(userId int64) error {
	if err := rpc_client.Call(addrs.UserAddr, "UserService.ReactivateUser", &userId, &struct{}{}); err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdatePassword(newUserPassword *model.NewUserPasswordDTO) error {
	user, err := us.GetUserById(newUserPassword.Id)
	if err != nil {
		return err
	}

	if err = validatePassword(user.Password, newUserPassword.OldPassword); err != nil {
		return err
	}

	user.Password = newUserPassword.NewPassword

	if err = rpc_client.Call(addrs.UserAddr, "UserService.UpdateUser", &user, &struct{}{}); err != nil {
		return err
	}

	return nil
}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
