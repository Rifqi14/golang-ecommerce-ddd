package routers

import (
	"github.com/Rifqi14/golang-ecommerce/app/handlers"
	v1 "github.com/Rifqi14/golang-ecommerce/routers/v1"
	"github.com/gofiber/fiber/v2"
)

func (boot Bootstrap) Router() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.DB,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	apiParentGroup := boot.App.Group("/api")
	apiParentGroup.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Api is working")
	})

	v1Routers := v1.V1Route{
		RouteGroup: apiParentGroup,
		Handler:    handlerType,
	}
	v1Routers.V1Router()
}
