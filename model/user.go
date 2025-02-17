package model

import (
	"database/sql"
	"time"
)

const (
	USER_STATUS_UNVERIFIED = "unverified"
	USER_STATUS_ACTIVE     = "active"

	GENDER_MALE   = "male"
	GENDER_FEMALE = "female"

	USER_ROLE_BASIC = "basic"
	USER_ROLE_ADMIN = "admin"
)

var (
	VALID_GENDERS = []string{GENDER_MALE, GENDER_FEMALE}
)

type (
	User struct {
		ID        int64        `db:"id"`         //
		CreatedAt time.Time    `db:"created_at"` //
		UpdatedAt time.Time    `db:"updated_at"` //
		DeletedAt sql.NullTime `db:"deleted_at"` //
		Guid      string       `db:"guid"`       //
		Email     string       `db:"email"`      //
		About     string       `db:"about"`      //
		Password  string       `db:"password"`   //
		Name      string       `db:"name"`       //
		Username  string       `db:"username"`   //
		PhotoUrl  string       `db:"photo_url"`  //
		UserRole  string       `db:"user_role"`  // Enum: basic, admin
	}
)
