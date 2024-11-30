package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/cron"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/handlers/ping_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_channel_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_video_handler"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/youtube_channel_repo"
	"github.com/umarkotak/ytkidd-api/repos/youtube_video_repo"
	"github.com/umarkotak/ytkidd-api/utils/log_formatter"
	"github.com/umarkotak/ytkidd-api/utils/log_hook"
	"github.com/umarkotak/ytkidd-api/utils/middlewares"
	"github.com/umarkotak/ytkidd-api/utils/ratelimit_lib"
	"github.com/umarkotak/ytkidd-api/utils/redis_lock"
	"github.com/umarkotak/ytkidd-api/utils/user_auth"
	"github.com/umarkotak/ytkidd-api/utils/word_censor_lib"
	"github.com/umarkotak/ytkidd-api/worker"
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&log_formatter.Formatter{})
	logrus.AddHook(&log_hook.LogrusHook{})

	config.Initialize()
	datastore.Initialize()

	args := os.Args[1:]

	if len(args) == 0 {
		runServer()

	} else if len(args) == 1 {
		switch args[0] {
		case "gen_base_key":
			user_auth.GenBaseKey()
		}

	} else if len(args) == 2 && args[0] == "migrate" {
		switch args[1] {
		case "up":
			datastore.MigrateUp()
		default:
			fmt.Println("Unknown migrate command. Use 'up' or 'seed'.")
		}

	} else {
		fmt.Println("Usage: go run .")
		fmt.Println("       go run . migrate up")
		fmt.Println("       go run . gen_base_key")
	}
}

func runServer() {
	logrus.Infof("starting ytkidd express backend")

	initializeDependencies()

	cron.Start()

	go worker.Run()

	go initializeHttpServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan, os.Interrupt, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM,
	)

	sig := <-sigChan
	logrus.Infof("received %v signal. graceful shutdown...", sig)

	// execute gracefull shutdown

	logrus.Infof("cleanup completed. exiting...")
}

func initializeDependencies() {
	var err error

	user_auth.Initialize(user_auth.UserAuth{
		JwtPrivateKey: config.Get().JxAuthJwtPrivateKey,
		JwtPublicKey:  config.Get().JxAuthJwtPublicKey,
		JweSecretKey:  config.Get().JxAuthJweSecretKey,
	})

	err = ratelimit_lib.Initialize(ratelimit_lib.RateLimiter{
		Prefix:    model.REDIS_PREFIX,
		RedisAddr: config.Get().RedisUrl,
		RedisPwd:  config.Get().RedisPassword,
	})
	if err != nil {
		logrus.WithContext(context.TODO()).Error(err)
	}

	err = redis_lock.Initialize(redis_lock.RedisLock{
		Prefix:    model.REDIS_PREFIX,
		RedisAddr: config.Get().RedisUrl,
		RedisPwd:  config.Get().RedisPassword,
	})
	if err != nil {
		logrus.WithContext(context.TODO()).Error(err)
	}

	err = worker.Initialize(config.Get().RedisUrl, config.Get().RedisPassword)
	if err != nil {
		logrus.WithContext(context.TODO()).Error(err)
	}

	err = cron.Initialize(config.Get().RedisUrl, config.Get().RedisPassword)
	if err != nil {
		logrus.WithContext(context.TODO()).Error(err)
	}

	youtube_channel_repo.Initialize()
	youtube_video_repo.Initialize()

	word_censor_lib.Initialize(word_censor_lib.WordCensorLib{
		Words: []string{"kucing", "anjing", "gajah"},
	})
}

func initializeHttpServer() {
	r := chi.NewRouter()

	r.Use(
		chiMiddleware.RequestID,                       //
		chiMiddleware.RealIP,                          //
		chiMiddleware.Recoverer,                       //
		middlewares.RequestLog,                        //
		middlewares.CommonCtx,                         // it will extract headers and put the value to common context
		middlewares.ReqRateLimit(1000, 1*time.Second), // max 100 request per second based on X-Device-Id
	)

	r.Route("/ytkidd/api", func(ri chi.Router) {
		// rDevInternal := ri.With(middlewares.InternalDevAuth)
		// rUserAuth := ri.With(middlewares.UserAuth)
		// rOptionalUserAuth := ri.With(middlewares.OptionalUserAuth)

		ri.Get("/ping", ping_handler.Ping)

		ri.Get("/youtube_videos", youtube_video_handler.GetYoutubeVideos)
		ri.Get("/youtube_video/{id}", youtube_video_handler.GetYoutubeVideoDetail)
		ri.Get("/youtube_channels", youtube_channel_handler.GetYoutubeChannels)

		ri.Post("/youtube/scrap_videos", youtube_handler.ScrapVideos)
	})

	port := fmt.Sprintf(":%s", config.Get().AppPort)
	logrus.Infof("running http server on port %s", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalf("fail to start http server on port %s", port)
	}
}
