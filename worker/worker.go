package worker

import (
	"context"
	"fmt"

	goWorker "github.com/digitalocean/go-workers2"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func defaultQueueProcessor(message *goWorker.Msg) error {
	ctx := context.TODO()

	ctx = context.WithValue(ctx, chiMiddleware.RequestIDKey, message.Jid())

	msgByte, _ := message.Args().Encode()
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"params": fmt.Sprintf("%+v", string(msgByte)),
	}).Infof("[WORKER][CLASS: %s]", message.Class())

	// TODO: process message based on class
	switch message.Class() {
	default:
	}

	return nil
}
