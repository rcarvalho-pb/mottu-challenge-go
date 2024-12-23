package services

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	rpc_client "github.com/rcarvalho-pb/mottu-authentication_service/internal/rpc/client"
)

type tokenService struct {
	addr string
}

func newTokenService(tokenServiceAddr string) *tokenService {
	return &tokenService{
		addr: tokenServiceAddr,
	}
}

func (ts *tokenService) getToken(user *dtos.UserDTO) (string, error) {
	userDto := struct {
		Id       int64
		Username string
	}{
		user.Id,
		user.Username,
	}

	var tokenString string
	if err := rpc_client.Call(ts.addr, "TokenService.GenerateToken", userDto, &tokenString); err != nil {
		return "", fmt.Errorf("error calling token service: %s\n", err)
	}

	return tokenString, nil
}
