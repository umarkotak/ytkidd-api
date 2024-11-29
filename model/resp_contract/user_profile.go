package resp_contract

import (
	"time"
)

type (
	MyProfile struct {
		Guid               string             `json:"guid"`                  //
		Phone              string             `json:"phone"`                 //
		Email              string             `json:"email"`                 //
		Name               string             `json:"name"`                  //
		Gender             string             `json:"gender"`                //
		DateOfBirth        string             `json:"date_of_birth"`         //
		Ethnic             string             `json:"ethnic"`                //
		Religion           string             `json:"religion"`              //
		PhoneVerified      bool               `json:"phone_verified"`        //
		IdVerified         bool               `json:"id_verified"`           //
		ProfilePictureUrl  string             `json:"profile_picture_url"`   //
		TotalLikes         int64              `json:"total_likes"`           // TODO: implement logic
		RealTimeTotalLikes int64              `json:"real_time_total_likes"` // TODO: implement logic
		TotalGiftGross     int64              `json:"total_gift_gross"`      // TODO: implement logic
		TotalGiftNett      int64              `json:"total_gift_nett"`       // TODO: implement logic
		TotalGiftCount     int64              `json:"total_gift_count"`      // TODO: implement logic
		MyGifts            []MyGiftData       `json:"my_gifts"`              //
		MyGifters          []MyGiftersData    `json:"my_gifters"`            //
		MyProfileGalleries []MyProfileGallery `json:"my_profile_galleries"`  //
	}

	MyGiftData struct {
		Title      string    `json:"title"`
		ImageUrl   string    `json:"image_url"`
		Amount     int64     `json:"amount"`
		ReceivedAt time.Time `json:"received_at"`
	}

	MyGiftersData struct {
		UserGuid          string    `json:"user_guid"`
		ProfilePictureUrl string    `json:"profile_picture_url"`
		Name              string    `json:"name"`
		ReceivedAt        time.Time `json:"received_at"`
	}

	MyPreference struct {
		MinAge    int64    `json:"min_age"`   //
		MaxAge    int64    `json:"max_age"`   //
		Ethnics   []string `json:"ethnics"`   //
		Religions []string `json:"religions"` //
		Tiers     []string `json:"tiers"`     //
	}

	MyProfileGallery struct {
		ID       int64  `json:"id"`
		Url      string `json:"url"`
		Priority int64  `json:"priority"`
	}

	UpdatePhoneNumber struct {
		OtpToken string `json:"otp_token"`
	}
)
