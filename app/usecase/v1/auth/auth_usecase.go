package v1_usecase

import (
	authRequest "github.com/Rifqi14/golang-ecommerce/app/domain/auth/request"
	ucinterface "github.com/Rifqi14/golang-ecommerce/app/domain/auth/usecase"
	authVm "github.com/Rifqi14/golang-ecommerce/app/domain/auth/view_models"
	userVm "github.com/Rifqi14/golang-ecommerce/app/domain/user/view_models"
	"github.com/Rifqi14/golang-ecommerce/app/usecase"
)

type AuthUsecase struct {
	*usecase.Contract
}

func NewAuthUsecase(contract *usecase.Contract) ucinterface.IAuthenticationUsecase {
	return &AuthUsecase{Contract: contract}
}

func (uc AuthUsecase) LoginByEmail(req *authRequest.LoginRequest) (res authVm.TokenVm, err error) {
	panic("implement me")
}

func (uc AuthUsecase) Register(req *authRequest.RegisterRequest) error {
	panic("implement me")
}

func (uc AuthUsecase) GetCurrentUser() (res userVm.UserVm, err error) {
	panic("implement me")
}

func (uc AuthUsecase) Logout() error {
	panic("implement me")
}

func (uc AuthUsecase) GenerateJWT(issuer, payload string) (res authVm.TokenVm, err error) {
	panic("implement me")
}
