package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/router"
)

func main() {
	config.Start()
	mux := router.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Addresses.BrokerAddr), mux))
}
