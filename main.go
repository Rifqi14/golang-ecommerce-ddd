package main

import (
	"fmt"
	"os"

	"github.com/Rifqi14/golang-ecommerce/app/domain"
	"github.com/Rifqi14/golang-ecommerce/config/functioncaller"
	"github.com/Rifqi14/golang-ecommerce/config/logruslogger"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	logFormat       = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}","user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
	validatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	// app_env := os.Getenv("APP_ENV")

	config, err := domain.LoadConfig()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-config")
	}

	db, err := config.DB.DB()
	defer db.Close()

	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	validatorDriver = validator.New()

	err := enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(validatorDriver, transID)
	if err != nil {
		fmt.Println(err)
	}

	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
