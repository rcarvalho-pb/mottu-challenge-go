package main

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LocationDTO struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	MotorcycleId  int64     `json:"motorcycle_id"`
	Price         float64   `json:"price"`
	Days          int64     `json:"days"`
	LocationDay   time.Time `json:"location_day"`
	DevolutionDay time.Time `json:"devolution_day"`
	Fine          float64   `json:"fine"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type MotorcycleDTO struct {
	Id        int64     `json:"id"`
	Year      int64     `json:"year"`
	Model     string    `json:"model"`
	Plate     string    `json:"plate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsLocated bool      `json:"is_located"`
	Active    bool      `json:"active"`
}

type UserDTO struct {
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           int       `json:"role"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birth_date"`
	CNPJ           string    `json:"cnpj"`
	CNH            string    `json:"cnh"`
	CNHType        string    `json:"cnh_type"`
	CNHFilePath    string    `json:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location"`
}

type Claims struct {
	UserID   int64
	Username string
	UserRole int
	jwt.RegisteredClaims
}

type UserDTO2 struct {
	Id       int64
	Username string
	Role     int
}

func Call[K, T any](port, service string, data K, entry *T) (err error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Printf("Erro ao conectar ao servidor na port [%s]: %s\n", port, err)
		return
	}
	defer client.Close()

	err = client.Call(service, &data, &entry)
	if err != nil {
		return
	}

	return
}

func main() {

	var user *UserDTO

	username := "rcarvalho"

	err := Call("12345", "UserService.GetUserByUsername", username, &user)
	if err != nil {
		fmt.Printf("erro ao chamar serviço: %s\n", err)
	} else {
		fmt.Println(user)
	}

	var users []*UserDTO

	// O serviço foi registrado como "UserService", não "RPCServer".
	err = Call("12345", "UserService.GetAllUsers", struct{}{}, &users)
	if err != nil {
		fmt.Println("Erro ao chamar o serviço:", err)
	} else {
		fmt.Println("Usuários encontrados:")
		for _, user := range users {
			fmt.Printf("- %+v\n", user)
		}
	}

	userDTO2 := UserDTO2{
		Id:       1,
		Username: "rcarvalho",
	}

	var tokenString string

	err = Call("12346", "TokenService.GenerateToken", userDTO2, &tokenString)
	if err != nil {
		fmt.Println("Erro ao chamar o serviço:", err)
	} else {
		fmt.Println(tokenString)
	}

	var claims Claims

	err = Call("12346", "TokenService.ValidateToken", tokenString, &claims)
	if err != nil {
		fmt.Println("Erro ao chamar serviço:", err)
	} else {
		fmt.Printf("%+v\n", claims)
	}

	var tokenString2 string

	if err = Call("12347", "AuthService.Authenticate", user, &tokenString2); err != nil {
		fmt.Println("Erro ao chamar o serviço:", err)
	} else {
		fmt.Println("Authenticate:", tokenString2)
	}

	var motorcycles []*MotorcycleDTO
	if err = Call("12348", "MotorcycleService.GetAllActiveMotorcycles", struct{}{}, &motorcycles); err != nil {
		fmt.Println("Erro ao chamar serviço:", err)
	} else {
		fmt.Println("morotcycles:")
		for _, m := range motorcycles {
			fmt.Printf("%+v\n", m)

		}
	}

	var locations []*LocationDTO
	if err = Call("12349", "LocationService.GetAllLocations", struct{}{}, &locations); err != nil {
		fmt.Println("Erro ao chamar serviço:", err)
	} else {
		fmt.Println("locations:")
		for _, l := range locations {
			fmt.Printf("%+v\n", l)
		}
	}
}
