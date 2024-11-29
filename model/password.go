package model

type (
	InitPasswordData struct {
		UserID            int64  `json:"user_id"`
		InitPasswordToken string `json:"otp_token"`
	}

	ResetPasswordData struct {
		UserID   int64  `json:"user_id"`
		OtpToken string `json:"otp_token"`
		OtpValue string `json:"otp_value"`
	}
)
