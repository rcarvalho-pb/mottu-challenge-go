package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"rcarvalho-pb/mottu-token_service/internal/application/dtos"
	"rcarvalho-pb/mottu-token_service/internal/application/services"
)

type RPCServer struct {
	tokenService *services.TokenService
}

func New(service *services.TokenService) *RPCServer {
	return &RPCServer{
		tokenService: service,
	}
}

func (r *RPCServer) RPCListen(port string) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("TokenService", r)
	if err != nil {
		fmt.Println(err)
	}

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			return err
		}

		go rpc.ServeConn(rpcConn)
	}
}

func (r *RPCServer) GenerateToken(dto dtos.UserDTO, reply *string) error {
	tokenString, err := r.tokenService.GenerateJWT(&dto)
	if err != nil {
		return err
	}

	*reply = tokenString

	return err
}

func (r *RPCServer) ValidateToken(tokenString string, reply *services.Claims) error {
	claims, err := r.tokenService.ValidateToken(tokenString)
	if err != nil {
		return fmt.Errorf("error validating token: %s", err)
	}

	*reply = *claims

	return nil
}
