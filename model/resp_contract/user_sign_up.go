package resp_contract

type (
	UserSignUp struct {
		OtpToken string `json:"otp_token"`
	}

	SignUpVerifyOtp struct {
		InitPasswordToken string `json:"init_password_token"`
	}
)
