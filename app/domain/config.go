package domain

import (
	"fmt"
	"log"
	"os"

	db "github.com/Rifqi14/golang-ecommerce/config/db_connection"
	"github.com/Rifqi14/golang-ecommerce/config/jwe"
	"github.com/Rifqi14/golang-ecommerce/config/jwt"
	"github.com/Rifqi14/golang-ecommerce/config/redis"
	"github.com/Rifqi14/golang-ecommerce/config/str"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	redisExt "github.com/go-redis/redis/v7"
	jwtFiber "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	JwtConfig     jwtFiber.Config
	Validator     *validator.Validate
	Redis         redis.RedisClient
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	// Setup DB Connection
	dbInfo := db.Connection{
		Driver:                  os.Getenv("DB_DRIVER"),
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USER"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		DBMaxConnections:        str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMaxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifetimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	res.DB, err = dbInfo.Conn()
	if err != nil {
		log.Fatal(err.Error())
	}

	// JWE Credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	// JWT Credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("JWT_TOKEN_SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("JWT_TOKEN_EXPIRED")),
		RefreshTokenSecret:  os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED")),
	}

	// JWT Config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(res.JwtCredential.TokenSecret),
		Claims:     &jwt.CustomClaims{},
	}

	// Setup Redis
	redisOption := &redisExt.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       str.StringToInt(os.Getenv("REDIS_DB")),
	}
	res.Redis = redis.RedisClient{
		Client: redisExt.NewClient(redisOption),
	}
	pong, err := res.Redis.Client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Redis ping status:"+pong, err)

	res.Validator = ValidatorDriver

	return res, err
}
