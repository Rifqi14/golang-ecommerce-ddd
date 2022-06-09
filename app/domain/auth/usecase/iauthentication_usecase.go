package auth

import (
	authRequest "github.com/Rifqi14/golang-ecommerce/app/domain/auth/request"
	auth "github.com/Rifqi14/golang-ecommerce/app/domain/auth/view_models"
	userVm "github.com/Rifqi14/golang-ecommerce/app/domain/user/view_models"
)

type IAuthenticationUsecase interface {
	LoginByEmail(req *authRequest.LoginRequest) (res auth.TokenVm, err error)

	Register(req *authRequest.RegisterRequest) error

	GetCurrentUser() (res userVm.UserVm, err error)

	Logout() (err error)

	GenerateJWT(issuer, payload string) (res auth.TokenVm, err error)
}
