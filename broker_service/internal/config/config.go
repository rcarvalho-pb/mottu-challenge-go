package config

import "os"

type Adresses struct {
	UserAddr       string
	AuthAddr       string
	LocationAddr   string
	MotorcycleAddr string
	BrokerAddr     string
}

var Addresses *Adresses

func Start() {
	Addresses.AuthAddr = os.Getenv("AUTH_SERVICE_ADDRESS")
	Addresses.UserAddr = os.Getenv("USER_SERVICE_ADDRESS")
	Addresses.LocationAddr = os.Getenv("LOCATION_SERVICE_ADDRESS")
	Addresses.MotorcycleAddr = os.Getenv("MOTORCYCLE_SERVICE_ADDRESS")
	Addresses.BrokerAddr = os.Getenv("BROKER_SERVICE_ADDRESS")
}
