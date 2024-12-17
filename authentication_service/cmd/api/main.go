package main

import (
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/services"
)

func main() {

	db := sqlite.GetDB()
	req := dtos.UserRequest{
		Username: "Andrey",
		Password: "123",
	}
	service := services.NewUserService(db)
	service.AuthUser(req)

}
