package user_auth

type (
	UserAuth struct {
		JwtPrivateKey string
		JwtPublicKey  string
		JweSecretKey  string
	}
)

var userAuth UserAuth

func Initialize(ua UserAuth) {
	userAuth = ua
}
