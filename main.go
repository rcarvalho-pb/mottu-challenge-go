package main

import (
	"fmt"
	"net/rpc"
	"reflect"
	"time"
)

type UserDTO struct {
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birth_date"`
	CNPJ           string    `json:"cnpj"`
	CNH            string    `json:"cnh"`
	CNHType        string    `json:"cnh_type"`
	CNHFilePath    string    `json:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location"`
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()

	var user *UserDTO

	err = client.Call("UserService.GetUserByUsername", "rcarvalho", user)
	if err != nil {
		fmt.Println("erro ao chamar serviço:", err)
	} else {
		fmt.Println(user)
	}

	var users []*UserDTO

	// O serviço foi registrado como "UserService", não "RPCServer".
	err = client.Call("UserService.GetAllUsers", struct{}{}, &users)
	if err != nil {
		fmt.Println("Erro ao chamar o serviço:", err)
	} else {
		fmt.Println("Usuários encontrados:")
		for _, user := range users {
			fmt.Printf("- %+v\n", user)
		}
	}

	fmt.Println(reflect.TypeOf(users))
}
