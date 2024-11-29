package encryption

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"

	kaJwt "github.com/kataras/jwt"

	"github.com/sirupsen/logrus"
)

func GenJwt(ctx context.Context, privateKey string, payload any) (string, error) {
	privateKeyByte, err := hex.DecodeString(privateKey)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}
	eddsaPrivateKey := ed25519.PrivateKey(privateKeyByte)

	jwtToken, err := kaJwt.Sign(kaJwt.EdDSA, eddsaPrivateKey, payload)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return string(jwtToken), nil
}

func DecodeJwt(ctx context.Context, publicKey, jwtToken string, dest any) error {
	publicKeyByte, err := hex.DecodeString(publicKey)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}
	eddsaPublicKey := ed25519.PublicKey(publicKeyByte)

	verifiedToken, err := kaJwt.Verify(kaJwt.EdDSA, eddsaPublicKey, []byte(jwtToken))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	err = verifiedToken.Claims(&dest)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
