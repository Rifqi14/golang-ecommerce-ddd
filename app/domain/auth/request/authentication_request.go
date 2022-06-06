package auth

type RegisterRequest struct {
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required"`
	RegisterBy      string `json:"register_by" form:"register_by"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ClaimOTPRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
	OTP   string `json:"otp" form:"otp" validate:"required"`
}
