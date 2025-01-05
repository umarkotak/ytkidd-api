package contract

type (
	UserSignUp struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name"`
		Username string `json:"username"`
		About    string `json:"about"`
	}

	UserSignIn struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
