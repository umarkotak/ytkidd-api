package encryption

import (
	"context"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(ctx context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(ctx context.Context, password, passwordHashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
