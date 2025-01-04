package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
)

type authService struct{}

func newAuthService() *authService {
	return &authService{}
}

func (as *authService) authenticate(user *model.UserDTO) (*string, error) {
	var (
		tokenString *string
		err         error
	)
	if user == nil {
		return nil, fmt.Errorf("user can't be null")
	}
	rpc_client.Call(addrs.AuthAddr, "AuthService.Authenticate", &user, &tokenString)
	return tokenString, err
}
