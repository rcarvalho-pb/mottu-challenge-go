package services

import (
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	rpc_client "github.com/rcarvalho-pb/mottu-authentication_service/internal/rpc/client"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	addr string
}

func newUserService(userServiceAddr string) *userService {
	return &userService{
		addr: userServiceAddr,
	}
}

func (us *userService) getUser(username string) (*dtos.UserDTO, error) {
	var userDto *dtos.UserDTO
	if err := rpc_client.Call(us.addr, "UserService.GetUserByUsername", username, &userDto); err != nil {
		return nil, err
	}

	return userDto, nil
}

func (us *userService) validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
