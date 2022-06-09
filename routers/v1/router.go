package v1

import (
	"github.com/Rifqi14/golang-ecommerce/app/handlers"
	v1 "github.com/Rifqi14/golang-ecommerce/routers/v1/auth"
	"github.com/gofiber/fiber/v2"
)

type V1Route struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (api V1Route) V1Router() {
	apiV1 := api.RouteGroup.Group("/v1")

	authRoute := v1.AuthRoute{
		RouteGroup: apiV1,
		Handler:    api.Handler,
	}
	authRoute.AuthRoute()
}
