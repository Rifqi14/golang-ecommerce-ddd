package handlers

import (
	"github.com/Rifqi14/golang-ecommerce/app/usecase"
	"github.com/Rifqi14/golang-ecommerce/config/jwe"
	"github.com/Rifqi14/golang-ecommerce/config/jwt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	App           *fiber.App
	UcContract    *usecase.Contract
	DB            *gorm.DB
	Validate      *validator.Validate
	Translator    ut.Translator
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
}
