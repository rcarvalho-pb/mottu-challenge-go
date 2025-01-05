package service

import (
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/dtos"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
)

type authService struct{}

func newAuthService() *authService {
	return &authService{}
}

func (as *authService) Authenticate(authRequest *dtos.AuthRequest) (tokenString *string, err error) {
	rpc_client.Call(addrs.AuthAddr, "AuthService.Authenticate", &authRequest, &tokenString)
	return tokenString, err
}
