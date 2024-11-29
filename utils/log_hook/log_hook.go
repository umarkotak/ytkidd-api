package log_hook

import (
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type LogrusHook struct{}

func (rih *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (rih *LogrusHook) Fire(e *logrus.Entry) error {
	if e.Context != nil {
		if e.Context.Value(middleware.RequestIDKey) != nil {
			reqId := fmt.Sprintf("%v", e.Context.Value(middleware.RequestIDKey))
			if reqId != "" {
				e.Data["request_id"] = reqId
			}
		}
	}
	return nil
}
