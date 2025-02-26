package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/services"
)

type RPCServer struct {
	services.Service
	port string
}

func New(port string, service services.Service) *RPCServer {
	return &RPCServer{
		port:    port,
		Service: service,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.port))
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("AuthService", r)
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

func (r *RPCServer) Authenticate(authRequest *dtos.AuthRequest, reply *string) error {
	tokenString, err := r.Service.AuthenticateUser(authRequest)
	if err != nil {
		return err
	}

	*reply = tokenString

	return err
}
