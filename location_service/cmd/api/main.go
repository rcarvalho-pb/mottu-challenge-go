package main

import (
	"log"
	"os"

	"github.com/rcarvalho-pb/mottu-location_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-location_service/internal/application/service"
	rpc_server "github.com/rcarvalho-pb/mottu-location_service/internal/rpc/server"
)

type Config struct {
	RPCServer *rpc_server.RPCServer
}

func main() {
	dbLocation := os.Getenv("DB_LOCATION")
	db := sqlite.GetDB(dbLocation)
	locationService := service.New(db)

	port := os.Getenv("LOCATION_SERVICE_ADDRESS")
	config := &Config{rpc_server.New(locationService, port)}

	log.Fatal(config.RPCServer.RPCListen())

}
