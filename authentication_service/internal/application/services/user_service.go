package services

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/rpc"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AuthUser(req dtos.UserRequest) (string, error) {
	user := new(dtos.UserDTO)
	if err := rpc.Call("UserService.GetUserByUsername", "localhost:12345", req.Username, user); err != nil {
		fmt.Println("error calling rpc server:", err)
	}

	if err := validatePassword(user.Password, req.Password); err != nil {
		fmt.Println(err)
	}

	userDto := struct {
		Id       int64
		Username string
	}{
		user.Id,
		user.Username,
	}

	var tokenString string
	if err := rpc.Call("TokenService", "localhost:12346", userDto, &tokenString); err != nil {
		return "", fmt.Errorf("error calling token service: %s\n", err)
	}

	return tokenString, nil

}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
