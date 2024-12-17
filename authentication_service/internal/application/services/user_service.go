package services

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/domain/model"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/rpc"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(repository model.UserRepository) *UserService {
	return &UserService{
		UserRepository: repository,
	}
}

func (us *UserService) AuthUser(req dtos.UserRequest) {
	user := new(dtos.UserDTO)
	if err := rpc.Call("UserService.GetUserByUsername", req.Username, user); err != nil {
		fmt.Println(err)
	}

	if err := validatePassword(user.Password, req.Password); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", user)

}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
