package worker

import (
	"context"

	goWorker "github.com/digitalocean/go-workers2"
	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_QUEUE = "default"
)

type (
	WorkerManager struct {
		Manager  *goWorker.Manager
		Producer *goWorker.Producer
	}
)

var workerManager WorkerManager

func Initialize(redisAddr, redisPwd string) error {
	manager, err := goWorker.NewManager(goWorker.Options{
		Namespace:          "JODOH_EXPRESS",
		ManagerDisplayName: "JODOH_EXPRESS",
		ServerAddr:         redisAddr,
		Password:           redisPwd,
		Database:           0,
		PoolSize:           30,
		ProcessID:          "1",
	})
	if err != nil {
		logrus.Error(err)
		return err
	}

	workerManager = WorkerManager{
		Manager:  manager,
		Producer: manager.Producer(),
	}

	return nil
}

func Run() {
	workerManager.Manager.AddWorker(DEFAULT_QUEUE, 10, defaultQueueProcessor)

	workerManager.Manager.Run()
}

func Enqueue(ctx context.Context, class string, args any, opts *goWorker.EnqueueOptions) error {
	if opts == nil {
		opts = &goWorker.EnqueueOptions{}
	}
	_, err := workerManager.Producer.EnqueueWithContext(ctx, DEFAULT_QUEUE, class, args, *opts)
	return err
}
