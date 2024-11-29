package datastore

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
)

type DataStore struct {
	Db    *sqlx.DB      // required
	Redis *redis.Client // required
}

var dataStore DataStore

func Initialize() error {
	db, err := sqlx.Connect("postgres", config.Get().DbURL)
	if err != nil {
		logrus.Error(err)
		return err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Get().RedisUrl,
		Password: config.Get().RedisPassword,
		DB:       0,
	})
	_, err = redisClient.Ping(context.TODO()).Result()
	if err != nil {
		logrus.Error(err)
		return err
	}

	dataStore = DataStore{
		Db:    db,
		Redis: redisClient,
	}

	return nil
}

func Get() DataStore {
	return dataStore
}
