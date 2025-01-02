package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-user_service/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(Repository model.UserRepository) *UserService {
	return &UserService{
		UserRepository: Repository,
	}
}

func (us *UserService) CreateUser(newUser *dtos.UserDTO) error {
	user := model.UserFromDTO(newUser)

	if err := us.UserRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetAllUsers() ([]*dtos.UserDTO, error) {
	users, err := us.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersDTO []*dtos.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, user.ToDTO())
	}

	return usersDTO, err
}

func (us *UserService) GetAllActiveUsers() ([]*dtos.UserDTO, error) {
	users, err := us.UserRepository.GetAllActiveUsers()
	if err != nil {
		return nil, err
	}

	var usersDTO []*dtos.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, user.ToDTO())
	}

	return usersDTO, err
}

func (us *UserService) GetUserById(id int64) (*dtos.UserDTO, error) {
	user, err := us.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (us *UserService) GetUserByUsername(username string) (*dtos.UserDTO, error) {
	user, err := us.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (us *UserService) NewUser(newUser *dtos.UserDTO) error {
	if err := validateNewUser(*newUser); err != nil {
		return err
	}

	if err := parameterizaNewUser(newUser); err != nil {
		return err
	}

	user := model.UserFromDTO(newUser)

	hashedPassowrd, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassowrd)

	if err := us.UserRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(userId int64, newUser *dtos.UserDTO) error {
	user, err := us.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	if err = parameterizaNewUser(newUser); err != nil {
		return err
	}

	if newUser.Name != "" {
		user.Name = newUser.Name
	}

	if newUser.Username != "" {
		user.Username = newUser.Username
	}

	if !newUser.BirthDate.IsZero() {
		user.BirthDate = newUser.BirthDate
	}

	if newUser.CNPJ != "" {
		user.CNPJ = newUser.CNPJ
	}

	if newUser.CNH != "" {
		user.CNH = newUser.CNH
	}

	if newUser.CNHType != "" {
		user.CNHType = newUser.CNHType
	}

	// if newUser.CNHFilePath != "" {
	// 	user.CNHFilePath = newUser.CNHFilePath
	// }

	if user.ActiveLocation != newUser.ActiveLocation {
		user.ActiveLocation = newUser.ActiveLocation
	}

	user.UpdatedAt = time.Now()

	if err = us.UserRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeactivateUserById(userId int64) error {
	user, err := us.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	user.Active = false
	user.UpdatedAt = time.Now()
	if err = us.UserRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) ReactivateUserById(userId int64) error {
	user, err := us.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	user.Active = true
	user.UpdatedAt = time.Now()
	if err = us.UserRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func parameterizaNewUser(newUser *dtos.UserDTO) error {
	var err error
	newUser.Name = strings.TrimSpace(newUser.Name)
	newUser.Username = strings.TrimSpace(newUser.Username)
	newUser.Password = strings.TrimSpace(newUser.Password)
	newUser.CNPJ = strings.TrimSpace(newUser.CNPJ)
	newUser.CNH = strings.TrimSpace(newUser.CNH)
	newUser.CNHType = strings.TrimSpace(newUser.CNHType)

	return err
}

func validateNewUser(user dtos.UserDTO) error {
	if user.Name == "" {
		return fmt.Errorf("User name can't be empty")
	}
	if user.Username == "" {
		return fmt.Errorf("User username can't be empty")
	}
	if user.Password == "" {
		return fmt.Errorf("User password can't be empty")
	}
	if user.BirthDate.IsZero() {
		return fmt.Errorf("User birthdate can't be empty")
	}
	if user.CNH == "" {
		return fmt.Errorf("User cnh can't be empty")
	}
	if user.CNPJ == "" {
		return fmt.Errorf("User cnpj can't be empty")
	}
	if user.CNHType == "" {
		return fmt.Errorf("User cnh type can't be empty")
	}
	return nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
