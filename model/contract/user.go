package contract

import (
	"mime/multipart"

	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils"
)

type (
	UpdateMyProfile struct {
		UserGuid           string         `json:"-"`        //
		ProfilePictureFile multipart.File `json:"-"`        //
		KtpFile            multipart.File `json:"-"`        //
		Religion           string         `json:"religion"` //
		// Phone              string         `json:"phone"`    // TODO: should separate update phone and the others

		PrefMinAge    int64    `json:"pref_min_age"`
		PrefMaxAge    int64    `json:"pref_max_age"`
		PrefReligions []string `json:"pref_religions"`
		PrefEthnics   []string `json:"pref_ethnics"`
		PrefTiers     []string `json:"pref_tiers"`
	}

	UpdateMyProfileGallery struct {
		UserGuid           string         `json:"-"` //
		ProfilePictureFile multipart.File `json:"-"` //
		ProfilePictureIdx  int64          `json:"-"` //
	}

	ChangeMyPassword struct {
		UserGuid                string `json:"-"`
		OldPassword             string `json:"old_password"`
		NewPassword             string `json:"new_password"`
		NewPasswordConfirmation string `json:"new_password_confirmation"`
	}

	UpdateMyPhoneNumber struct {
		UserGuid       string `json:"-"`
		NewPhoneNumber string `json:"new_phone_number"`
	}

	UpdateMyPhoneNumberConfirmation struct {
		UserGuid string `json:"-"`
		OtpToken string `json:"otp_token"`
		Otp      string `json:"otp"`
	}

	DeleteUser struct {
		UserGuid string `json:"-"`
		Password string `json:"password"`
	}

	GetForDiscovery struct {
		UserGuid string // from user auth context

		UserID        int64    // from user db
		Age           int64    // from user db
		Gender        string   // from guest-gender or user db
		Religion      string   // from user db
		Ethnic        string   // from user db
		Tier          string   // from user db
		Name          string   // from query params
		PrefMinAge    int64    // from query params
		PrefMaxAge    int64    // from query params
		PrefReligions []string // from query params
		PrefEthnics   []string // from query params
		PrefTiers     []string // from query params

		Limit  int64 //  from query params
		Page   int64 //  from query params
		Offset int64 //  calculated from limit and page

		UpTierForTeaser bool    // set when call func
		ExcludeUserIDs  []int64 // set when up tier for teaser
	}
)

func (m *GetForDiscovery) ValidateAndFormat() error {
	if m.Limit < 1 {
		m.Limit = 1
	}

	if m.Page < 1 {
		m.Page = 1
	}

	if m.PrefMinAge < 17 {
		m.PrefMinAge = 17
	}

	if m.PrefMaxAge < 17 {
		m.PrefMaxAge = 130
	}

	if m.PrefMinAge > m.PrefMaxAge {
		return model.ErrBadRequest
	}

	m.Offset = (m.Page - 1) * m.Limit

	if m.UserGuid == "" {
		if !utils.SliceStringContain(model.VALID_GENDERS, m.Gender) {
			return model.ErrMissingGenderParams
		}
	}

	m.Tier = model.TIER_FREE

	return nil
}
