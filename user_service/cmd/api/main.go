package main

import (
	"fmt"
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
	"github.com/rcarvalho-pb/mottu-user_service/internal/domain/model"
)

func main() {
	db := sqlite.GetDB()
	userService := services.NewUserService(db)
	user := model.User{
		Username:       "rcarvalho",
		Password:       "123",
		Name:           "Ramon",
		BirthDate:      time.Time{},
		CNPJ:           "123123123",
		CNH:            "123123123123",
		CNHType:        "B",
		ActiveLocation: false,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		Active:         true,
	}

	if err := userService.UserRepository.CreateUser(&user); err != nil {
		fmt.Println(err)
	}

	saved, _ := userService.UserRepository.GetUserByUsername("rcarvalho")

	fmt.Printf("%+v\n", saved)
}
