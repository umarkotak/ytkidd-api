package cron

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	redislock "github.com/go-co-op/gocron-redis-lock/v2"
)

type (
	CronManager struct {
		Scheduler gocron.Scheduler
	}
)

var cronManager CronManager

func Initialize(redisAddr, redisPwd string) error {
	redisOptions := &redis.Options{
		Addr:     redisAddr,
		Password: redisPwd,
	}
	redisClient := redis.NewClient(redisOptions)

	locker, err := redislock.NewRedisLocker(redisClient, redislock.WithTries(1))
	if err != nil {
		logrus.Error(err)
		return err
	}

	scheduler, err := gocron.NewScheduler(gocron.WithDistributedLocker(locker))
	if err != nil {
		logrus.Error(err)
		return err
	}

	cronManager = CronManager{
		Scheduler: scheduler,
	}

	return nil
}

func Start() error {
	_, err := cronManager.Scheduler.NewJob(
		gocron.DurationJob(500*time.Millisecond),
		gocron.NewTask(func() {
			// logrus.Infof("CRON EXECUTED - Every 500ms")
		}),
	)
	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = cronManager.Scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(1, 0, 0))),
		gocron.NewTask(func() {
			// logrus.Infof("CRON EXECUTED - Daily at 01:00:00")
			DailyUserPlanReminder()
		}),
	)
	if err != nil {
		logrus.Error(err)
		return err
	}

	cronManager.Scheduler.Start()

	return nil
}
