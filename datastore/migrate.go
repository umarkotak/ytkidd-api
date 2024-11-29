package datastore

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
)

func MigrateUp() error {
	m, err := migrate.New("file://db/migrations", config.Get().DbURL)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logrus.Error(err)
		return err
	}

	return nil
}
