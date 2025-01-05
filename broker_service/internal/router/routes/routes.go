package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/controllers"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

var controller *controllers.Controller

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
	Admin          bool
}

func StartMux(mux *http.ServeMux) *http.ServeMux {
	controller = controllers.NewController(service.New(config.Addresses))
	var routes []*Route
	routes = append(routes, AuthRoutes...)
	routes = append(routes, UserRoutes...)
	routes = append(routes, LocationRoutes...)
	routes = append(routes, MotorcycleRoutes...)
	for _, route := range routes {
		if route.Authentication {
			mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Uri), route.Function)
		} else {

		}
	}

	return mux
}
