package main

import (
	"log"
	"os"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
	"github.com/rcarvalho-pb/mottu-user_service/internal/rpc"
)

type Config struct {
	RPCServer *rpc.RPCServer
}

func main() {
	db := sqlite.GetDB()
	userService := services.NewUserService(db)

	port := os.Getenv("USER_SERVICE_ADDRESS")
	config := &Config{rpc.New(*userService, port)}

	log.Fatal(config.RPCServer.RPCListen())
}
