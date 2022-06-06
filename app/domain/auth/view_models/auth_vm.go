package auth

type TokenVm struct {
	Token                 string `json:"token"`
	TokenExpiredAt        int64  `json:"token_expired_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiredAt int64  `json:"refresh_token_expired_at"`
}

type AuthVm struct {
	Token TokenVm `json:"token"`
}

func NewAuthVm() AuthVm {
	return AuthVm{}
}

func (vm AuthVm) BuildToken(token, refresh_token string, expired_at, refresh_expired_at int64) TokenVm {
	return TokenVm{
		Token:                 token,
		TokenExpiredAt:        expired_at,
		RefreshToken:          refresh_token,
		RefreshTokenExpiredAt: refresh_expired_at,
	}
}
