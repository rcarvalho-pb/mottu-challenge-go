package main

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
)

func main() {
	db := sqlite.GetDB()
	userService := services.NewUserService(db)
	// user := model.User{
	// 	Username:  "rcarvalho",
	// 	Password:  "123",
	// 	Name:      "Ramon223",
	// 	BirthDate: time.Now(),
	// 	CNPJ:      "123123122232",
	// 	CNH:       "123123131232123212",
	// 	CNHType:   "B",
	// }
	//
	// if err := userService.NewUser(user.ToDTO()); err != nil {
	// 	fmt.Println(err)
	// }

	users, err := userService.GetUsersByUsername("rcarv")
	if err != nil {
		fmt.Println(err)
	}

	for i, user := range users {
		fmt.Printf("%d - %+v\n", i, user)
	}
}
