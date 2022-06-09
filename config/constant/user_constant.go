package user_constant

const (
	DraftUserStatus     = "draft"
	ActiveUserStatus    = "active"
	NonActiveUserStatus = "nonactive"

	ByEmail    = "email"
	ByGoogle   = "google"
	ByFacebook = "facebook"

	PrefixRedisKeyOtp   = "otp-"
	PrefixRedisKeyToken = "token-"
)
