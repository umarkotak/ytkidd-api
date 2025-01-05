package user_auth

import (
	"github.com/redis/go-redis/v9"
)

type (
	UserAuth struct {
		JwtPrivateKey string
		JwtPublicKey  string
		JweSecretKey  string
		Redis         *redis.Client
	}
)

var userAuth UserAuth

func Initialize(ua UserAuth) {
	userAuth = ua
}
