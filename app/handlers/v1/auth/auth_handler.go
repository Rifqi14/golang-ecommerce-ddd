package v1_handler

import (
	hinterface "github.com/Rifqi14/golang-ecommerce/app/domain/auth/handler"
	"github.com/Rifqi14/golang-ecommerce/app/handlers"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	handlers.Handler
}

func NewAuthHandler(handler handlers.Handler) hinterface.IAuthenticationHandler {
	return &AuthHandler{Handler: handler}
}

func (handler AuthHandler) ClaimOTP(ctx *fiber.Ctx) error {
	return nil
}

func (handler AuthHandler) GetCurrentUser(ctx *fiber.Ctx) error {
	return nil
}

func (handler AuthHandler) Login(ctx *fiber.Ctx) error {
	return nil
}

func (handler AuthHandler) Register(ctx *fiber.Ctx) error {
	return nil
}

func (handler AuthHandler) RequestOTP(ctx *fiber.Ctx) error {
	return nil
}

func (handler AuthHandler) Logout() error {
	return nil
}
