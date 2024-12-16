package main

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
)

func main() {
	db := sqlite.GetDB()
	userService := services.NewUserService(db)

	users, err := userService.GetAllActiveUsers()
	if err != nil {
		fmt.Println(err)
	}

	for _, u := range users {
		fmt.Println(u)
	}

	fmt.Println(users)
}
