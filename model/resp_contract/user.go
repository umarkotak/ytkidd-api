package resp_contract

type (
	UserSignUp struct{}

	UserSignIn struct {
		AccessToken string `json:"access_token"`
	}
)
