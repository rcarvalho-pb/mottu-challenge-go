package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

func newUserService() *userService {
	return &userService{}
}

func (us *userService) CreateUser(newUser *model.UserDTO) error {
	if newUser == nil {
		return fmt.Errorf("user can't be null")
	}
	if err := rpc_client.Call(addrs.UserAddr, "UserService.CreateUser", &newUser, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *userService) UpdateUser(user *model.UserDTO) error {
	if user == nil {
		return fmt.Errorf("user can't be null")
	}
	if err := rpc_client.Call(addrs.UserAddr, "UserService.UpdateUser", &user, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *userService) GetAllUsers() ([]*model.UserDTO, error) {
	var users []*model.UserDTO
	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetAllUsers", &users, &struct{}{}); err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) GetAllActiveUsers() ([]*model.UserDTO, error) {
	var users []*model.UserDTO
	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetAllActiveUsers", &users, &struct{}{}); err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) GetUserById(userId int64) (*model.UserDTO, error) {
	var user *model.UserDTO
	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetUserById", &user, &userId); err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserByUsername(username string) (*model.UserDTO, error) {
	var user *model.UserDTO
	if err := rpc_client.Call(addrs.UserAddr, "UserService.GetUserByUsername", &username, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) DeactivateUser(userId int64) error {
	if err := rpc_client.Call(addrs.UserAddr, "UserService.DeactivateUser", &userId, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *userService) ReactivateUser(userId int64) error {
	if err := rpc_client.Call(addrs.UserAddr, "UserService.ReactivateUser", &userId, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *userService) UpdatePassword(user *model.UserDTO, newPassword string) error {
	if err := validatePassword(user.Password, newPassword); err != nil {
		return err
	}
	user.Password = newPassword
	if err := rpc_client.Call(addrs.UserAddr, "UserService.UpdateUser", &user, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
