package auth

import "github.com/gofiber/fiber/v2"

type IAuthenticationHandler interface {
	Login(ctx *fiber.Ctx) error

	LoginByEmail(ctx *fiber.Ctx) error

	Register(ctx *fiber.Ctx) error

	RequestOTP(ctx *fiber.Ctx) error

	GetCurrentUser(ctx *fiber.Ctx) error

	ClaimOTP(ctx *fiber.Ctx) error
}
