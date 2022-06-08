package usecase

import (
	"github.com/Rifqi14/golang-ecommerce/config/jwe"
	"github.com/Rifqi14/golang-ecommerce/config/jwt"
	"github.com/Rifqi14/golang-ecommerce/config/redis"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Contract struct {
	UserID        string
	App           *fiber.App
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	Validate      *validator.Validate
	Translator    ut.Translator
	Redis         redis.RedisClient
}
