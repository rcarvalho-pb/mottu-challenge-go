package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
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
