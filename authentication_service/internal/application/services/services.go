package services

import "github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"

type Service struct {
	*userService
	*tokenService
}

func New(userServiceAddr, tokenServiceAddr string) *Service {
	return &Service{newUserService(userServiceAddr), newTokenService(tokenServiceAddr)}
}

func (s *Service) AuthenticateUser(request dtos.UserRequest) (string, error) {
	user, err := s.getUser(request.Username)
	if err != nil {
		return "", err
	}

	if err = s.validatePassword(user.Password, request.Password); err != nil {
		return "", err
	}

	tokenString, err := s.GetToken(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
