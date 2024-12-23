package main

import (
	"log"
	"os"

	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/services"
	rpc_server "github.com/rcarvalho-pb/mottu-authentication_service/internal/rpc/server"
)

func main() {

	serverPort := os.Getenv("AUTH_SERVICE_ADDRESS")
	userServiceAddr := os.Getenv("USER_SERVICE_ADDRESS")
	tokenServiceAddr := os.Getenv("TOKEN_SERVICE_ADDRESS")

	service := services.New(userServiceAddr, tokenServiceAddr)
	r := rpc_server.New(serverPort, *service)

	log.Fatal(r.RPCListen())
}
