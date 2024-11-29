package resp_contract

type (
	Auth struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	InitResetPassword struct {
		OtpToken string `json:"otp_token"`
	}

	ResetPassword struct {
		ResetPasswordToken string `json:"reset_password_token"`
	}
)
