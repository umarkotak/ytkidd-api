package user_auth

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"

	kaJwt "github.com/kataras/jwt"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/utils/random"
	jose "gopkg.in/square/go-jose.v2"
)

func GenToken(ctx context.Context, userAuthPayload Payload) (string, error) {
	jwtToken, err := genJwtToken(ctx, userAuthPayload)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	accessToken, err := genJweToken(ctx, jwtToken)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return accessToken, nil
}

func genJwtToken(ctx context.Context, userAuthPayload Payload) ([]byte, error) {
	privateKey, err := hex.DecodeString(userAuth.JwtPrivateKey)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return []byte{}, err
	}
	eddsaPrivateKey := ed25519.PrivateKey(privateKey)

	jwtToken, err := kaJwt.Sign(kaJwt.EdDSA, eddsaPrivateKey, userAuthPayload)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return []byte{}, err
	}

	return jwtToken, nil
}

func genJweToken(ctx context.Context, jwtToken []byte) (string, error) {
	recepient := jose.Recipient{
		Algorithm:  jose.DIRECT,
		Key:        []byte(userAuth.JweSecretKey),
		PBES2Count: 2048,
		PBES2Salt:  random.GenBytes(2048),
	}

	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		recepient,
		&jose.EncrypterOptions{
			Compression: jose.DEFLATE,
		},
	)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	encryptedObject, err := encrypter.Encrypt(jwtToken)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	jweToken, err := encryptedObject.CompactSerialize()
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return jweToken, nil
}
