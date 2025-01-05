package user_auth

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"

	kaJwt "github.com/kataras/jwt"
	"github.com/sirupsen/logrus"
	jose "gopkg.in/square/go-jose.v2"
)

func VerifyAccessToken(ctx context.Context, accessToken string, opts VerifyOpts) (Payload, string, error) {
	encryptedObject, err := jose.ParseEncrypted(accessToken)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return Payload{}, "Parse Failed", err
	}

	jwtToken, err := encryptedObject.Decrypt([]byte(userAuth.JweSecretKey))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return Payload{}, "Invalid Token", err
	}

	publicKey, err := hex.DecodeString(userAuth.JwtPublicKey)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return Payload{}, "Invalid Token", err
	}
	eddsaPublicKey := ed25519.PublicKey(publicKey)

	verifiedToken, err := kaJwt.Verify(kaJwt.EdDSA, eddsaPublicKey, jwtToken)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		if err == kaJwt.ErrExpired {
			return Payload{}, "Token Expired", err
		}
		return Payload{}, "Verification Failed", err
	}

	var payload Payload
	err = verifiedToken.Claims(&payload)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return Payload{}, "Claim Failed", err
	}

	// TODO: session validation to redis

	return payload, "", nil
}
