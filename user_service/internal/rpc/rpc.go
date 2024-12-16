package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

type RPCServer struct{}

func (r *RPCServer) RPCListen() error {
	PORT := "5001"
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil {
		return err
	}

	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			return err
		}

		go rpc.ServeConn(rpcConn)
	}
}
