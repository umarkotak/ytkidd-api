package model

const (
	USER_STATUS_UNVERIFIED = "unverified"
	USER_STATUS_ACTIVE     = "active"

	GENDER_MALE   = "male"
	GENDER_FEMALE = "female"
)

var (
	VALID_GENDERS = []string{GENDER_MALE, GENDER_FEMALE}
)

type (
	User struct {
		ID          string `db:"id"`
		CreatedAt   string `db:"created_at"`
		UpdatedAt   string `db:"updated_at"`
		DeletedAt   string `db:"deleted_at"`
		Guid        string `db:"guid"`
		Phone       string `db:"phone"`
		Email       string `db:"email"`
		About       string `db:"about"`
		Password    string `db:"password"`
		Name        string `db:"name"`
		Gender      string `db:"gender"`
		DateOfBirth string `db:"date_of_birth"`
	}
)
