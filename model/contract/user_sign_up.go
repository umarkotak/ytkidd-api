package contract

import (
	"mime/multipart"
	"time"
)

type (
	UserSignUp struct {
		ProfilePictureFile multipart.File `json:"-"`
		KtpFile            multipart.File `json:"-"`
		Name               string         `json:"name"`
		Gender             string         `json:"gender"`
		DateOfBirth        time.Time      `json:"date_of_birth"`
		Ethnic             string         `json:"ethnic"`
		Religion           string         `json:"religion"`
		Phone              string         `json:"phone"`
		Email              string         `json:"email"`

		PrefMinAge    int64    `json:"pref_min_age"`
		PrefMaxAge    int64    `json:"pref_max_age"`
		PrefReligions []string `json:"pref_religions"`
		PrefEthnics   []string `json:"pref_ethnics"`
		PrefTiers     []string `json:"pref_tiers"`
	}

	UserSignUpVerifyOtp struct {
		OtpToken string `json:"otp_token"`
		Otp      string `json:"otp"`
	}

	UserSignUpSetInitialPassword struct {
		InitPassToken        string `json:"init_password_token"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}
)

func (m *UserSignUp) Reformat() error {
	return nil
}

func (m *UserSignUp) Validate() error {
	return nil
}
