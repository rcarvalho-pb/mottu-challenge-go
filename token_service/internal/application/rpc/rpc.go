package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"rcarvalho-pb/mottu-token_service/internal/application/dtos"
	"rcarvalho-pb/mottu-token_service/internal/application/services"
)

type RPCServer struct {
	*services.TokenService
	port string
}

func New(port string, service *services.TokenService) *RPCServer {
	return &RPCServer{
		TokenService: service,
		port:         port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.port))
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
	tokenString, err := r.GenerateJWT(&dto)
	if err != nil {
		return err
	}

	*reply = tokenString

	return err
}

func (r *RPCServer) ValidateToken(tokenString string, reply *dtos.Claims) error {
	claims, err := r.GetClaims(tokenString)
	if err != nil {
		return fmt.Errorf("error validating token: %s", err)
	}

	*reply = *claims

	return nil
}
