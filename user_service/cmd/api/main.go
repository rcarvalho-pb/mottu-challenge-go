package main

import (
	"fmt"
	"os"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
	"github.com/rcarvalho-pb/mottu-user_service/internal/rpc"
)

func main() {
	db := sqlite.GetDB()
	userService := services.NewUserService(db)

	r := rpc.New(*userService)

	port := os.Getenv("PORT")
	fmt.Println("PORT:", port)
	fmt.Println("Starting user service")
	r.RPCListen("12345")
}
