package services

import (
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/domain/model"
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

func (us *UserService) AuthUser(req dtos.UserRequest) (*model.User, error) {
	user, err := us.UserRepository.FindUserByUsername(req)
	if err != nil {
		return nil, err
	}

	// if err = validatePassword(req.Password, user.Password); err != nil {
	// 	return nil, fmt.Errorf("inválid user or password")
	// }

	return user, nil
}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
