package contract

type (
	InitResetPassword struct {
		Phone string `json:"phone"`
	}

	ResetPassword struct {
		OtpToken string `json:"otp_token"`
		Otp      string `json:"otp"`
	}

	ConfirmResetPassword struct {
		ResetPasswordToken      string `json:"reset_password_token"`
		NewPassword             string `json:"new_password"`
		NewPasswordConfirmation string `json:"new_password_confirmation"`
	}
)
