package main

import (
	"fmt"
	"rcarvalho-pb/mottu-token_service/internal/application/rpc"
	"rcarvalho-pb/mottu-token_service/internal/application/services"
)

func main() {

	tokenService := services.NewTokenService()

	r := rpc.New(tokenService)

	fmt.Println("Starting user service")
	r.RPCListen("12346")
}
