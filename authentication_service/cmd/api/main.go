package main

import (
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/services"
)

func main() {

	req := dtos.UserRequest{
		Username: "rcarvalho",
		Password: "123",
	}
	service := services.NewUserService()
	service.AuthUser(req)

}
