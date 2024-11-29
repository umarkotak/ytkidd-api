package resp_contract

type (
	UserForDiscoveryData struct {
		Guid               string `json:"guid"`
		Name               string `json:"name"`
		Gender             string `json:"gender"`
		Religion           string `json:"religion"`
		Ethnic             string `json:"ethnic"`
		ProfilePictureUrl  string `json:"profile_picture_tc_object_id"`
		Age                int64  `json:"age"`
		TotalLikes         int64  `json:"total_likes"`
		RealTimeTotalLikes int64  `json:"real_time_total_likes"`
		TierPlanCodename   string `json:"tier_plan_codename"`
	}

	UserProfileData struct {
		Guid               string             `json:"guid"`                  //
		Phone              string             `json:"phone"`                 //
		Email              string             `json:"email"`                 //
		Name               string             `json:"name"`                  //
		Gender             string             `json:"gender"`                //
		DateOfBirth        string             `json:"date_of_birth"`         //
		Age                int64              `json:"age"`                   //
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
		Gifts              []MyGiftData       `json:"gifts"`                 //
		Gifters            []MyGiftersData    `json:"gifters"`               //
		MyProfileGalleries []MyProfileGallery `json:"my_profile_galleries"`  //
		Tier               string             `json:"tier"`                  //
	}
)
