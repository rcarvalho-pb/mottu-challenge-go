package rpc

import (
	"fmt"
	"net/rpc"
)

type RPCServer struct {
}

func Call[T any](service string, address string, args any, reply *T) error {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return err
	}
	defer client.Close()

	// O serviço foi registrado como "UserService", não "RPCServer".
	err = client.Call(service, args, reply)
	if err != nil {
		fmt.Println("Erro ao chamar o serviço:", err)
	}

	return nil
}
