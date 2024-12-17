package main

import (
	"fmt"
	"net/rpc"
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
		fmt.Println("error connecting to the server:", err)
	}

	defer client.Close()

	var user []*UserDTO

	err = client.Call("RPCServer.GetAllUsers", struct{}{}, &user)
	if err != nil {
		fmt.Println("error calling service", err)
	} else {
		for _, u := range user {
			fmt.Printf("%+v\n", u)
		}
	}
}
