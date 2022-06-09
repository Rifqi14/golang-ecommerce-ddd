package v1

import (
	"github.com/Rifqi14/golang-ecommerce/app/handlers"
	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AuthRoute) AuthRoute() {
}
