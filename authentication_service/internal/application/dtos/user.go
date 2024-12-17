package dtos

import "time"

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

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
