package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-location_service/internal/application/service"
)

type RPCServer struct {
	locationService *service.LocationService
	Port            string
}

func New(locationService *service.LocationService, port string) *RPCServer {
	return &RPCServer{
		locationService: locationService,
		Port:            port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("UserService", r)
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
