package datastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
)

type DataStore struct {
	Db              *sqlx.DB          // required
	Redis           *redis.Client     // required
	R2Client        *s3.Client        //
	R2Manager       *manager.Uploader //
	R2PresignClient *s3.PresignClient //
	R2PublicDomain  string            //
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

	r2Config := aws.Config{
		Region:      "auto",
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.Get().R2AccessKeyId, config.Get().R2AccessKeySecret, "")),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					URL:               config.Get().R2StorageEndpoint,
					HostnameImmutable: true,
				}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		}),
	}

	r2Client := s3.NewFromConfig(r2Config, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	dataStore = DataStore{
		Db:              db,
		Redis:           redisClient,
		R2Client:        r2Client,
		R2Manager:       manager.NewUploader(r2Client),
		R2PresignClient: s3.NewPresignClient(r2Client),
		R2PublicDomain:  config.Get().R2PublicDomain,
	}

	return nil
}

func Get() DataStore {
	return dataStore
}
