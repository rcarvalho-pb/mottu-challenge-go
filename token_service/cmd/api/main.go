package main

import (
	"log"
	"os"
	"rcarvalho-pb/mottu-token_service/internal/application/rpc"
	"rcarvalho-pb/mottu-token_service/internal/application/services"
)

func main() {

	tokenServiceAddr := os.Getenv("TOKEN_SERVICE_ADDRESS")

	tokenService := services.NewTokenService()

	r := rpc.New(tokenServiceAddr, tokenService)

	log.Fatal(r.RPCListen())
}
