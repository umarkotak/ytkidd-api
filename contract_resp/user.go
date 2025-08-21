package contract_resp

type (
	UserSignUp struct{}

	UserSignIn struct {
		AccessToken string `json:"access_token"`
	}
)
