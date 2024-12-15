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
		Username:    "rcarvalho",
		Password:    "123",
		Name:        "Ramon223",
		BirthDate:   time.Now(),
		CNPJ:        "123123122232",
		CNH:         "123123131232123212",
		CNHType:     "B",
		CNHFilePath: "TESTE",
	}

	if err := userService.NewUser(user.ToDTO()); err != nil {
		fmt.Println(err)
	}

	saved, err := userService.GetUsersByUsername("rcarv")
	if err != nil {
		fmt.Println(err)
	}

	for i, u := range saved {
		fmt.Printf("%d - %s\n", i, u.Name)
	}
}
