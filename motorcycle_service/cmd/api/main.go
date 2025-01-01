package main

import (
	"log"
	"os"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/rpc"
)

type Config struct {
	RPCServer *rpc.RPCServer
}

func main() {
	dbLocation := os.Getenv("DB_LOCATION")
	db := sqlite.GetDB(dbLocation)
	motorcycleService := service.New(db)

	port := os.Getenv("USER_SERVICE_ADDRESS")
	config := &Config{rpc.New(*motorcycleService, port)}

	log.Fatal(config.RPCServer.RPCListen())
}
