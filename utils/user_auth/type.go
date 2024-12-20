package user_auth

type (
	Payload struct {
		JWTID          string `json:"jti"`
		SID            string `json:"sid"`
		Issuer         string `json:"iss"`
		IssuedAt       int64  `json:"iat"`
		ExpirationTime int64  `json:"exp"`
		GUID           string `json:"guid"`
		Name           string `json:"name"`
	}
)
