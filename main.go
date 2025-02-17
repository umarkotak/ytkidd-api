package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/cron"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/handlers/ai_handler"
	"github.com/umarkotak/ytkidd-api/handlers/book_handler"
	"github.com/umarkotak/ytkidd-api/handlers/comfy_ui_handler"
	"github.com/umarkotak/ytkidd-api/handlers/file_bucket_handler.go"
	"github.com/umarkotak/ytkidd-api/handlers/ping_handler"
	"github.com/umarkotak/ytkidd-api/handlers/user_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_channel_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_handler"
	"github.com/umarkotak/ytkidd-api/handlers/youtube_video_handler"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/repos/google_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
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
	logrus.Infof("starting ytkidd backend")

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
		JwtPrivateKey: config.Get().CkAuthJwtPrivateKey,
		JwtPublicKey:  config.Get().CkAuthJwtPublicKey,
		JweSecretKey:  config.Get().CkAuthJweSecretKey,
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

	// Repositories
	youtube_channel_repo.Initialize()
	youtube_video_repo.Initialize()
	book_repo.Initialize()
	book_content_repo.Initialize()
	file_bucket_repo.Initialize()
	user_repo.Initialize()
	google_repo.Initialize()

	word_censor_lib.Initialize(word_censor_lib.WordCensorLib{
		Words: []string{"kucing", "anjing", "gajah"},
	})
}

func initializeHttpServer() {
	r := chi.NewRouter()

	r.Use(
		chiMiddleware.RequestID, //
		chiMiddleware.RealIP,    //
		chiMiddleware.Recoverer, //
		middlewares.Cors,        //
		middlewares.RequestLog,  //
		middlewares.CommonCtx,   // it will extract headers and put the value to common context
		// middlewares.ReqRateLimit(1000, 1*time.Second), // max 100 request per second based on X-Device-Id
	)

	r.Route("/ytkidd/api", func(ri chi.Router) {
		// rDevInternal := ri.With(middlewares.InternalDevAuth)
		rUserAuth := ri.With(middlewares.UserAuth)
		// rOptionalUserAuth := ri.With(middlewares.OptionalUserAuth)

		ri.Get("/ping", ping_handler.Ping)

		ri.Get("/youtube_videos", youtube_video_handler.GetYoutubeVideos)
		ri.Get("/youtube_video/{id}", youtube_video_handler.GetYoutubeVideoDetail)
		ri.Get("/youtube_channels", youtube_channel_handler.GetYoutubeChannels)
		ri.Get("/youtube_channel/{id}", youtube_channel_handler.GetYoutubeChannelDetail)

		ri.Post("/youtube/scrap_videos", youtube_handler.ScrapVideos)

		ri.Post("/books/insert_from_pdf", book_handler.InsertFromPdf)
		ri.Post("/books/insert_from_pdf_urls", book_handler.InsertFromPdfUrls)
		ri.Get("/books", book_handler.GetBooks)
		ri.Get("/book/{id}", book_handler.GetBookDetail)
		ri.Delete("/book/{id}", book_handler.DeleteBook)

		ri.Get("/comfy_ui/output", comfy_ui_handler.Gallery)
		ri.Post("/ai/chat", ai_handler.Chat)

		ri.Post("/user/sign_in", user_handler.SignIn)
		rUserAuth.Get("/user/check_auth", user_handler.CheckAuth)
		rUserAuth.Get("/user/profile", user_handler.MyProfile)
		rUserAuth.Get("/user/subscription", ping_handler.ToDo)

		rUserAuth.Post("/order/new", ping_handler.ToDo)
		rUserAuth.Post("/order/pay", ping_handler.ToDo)

		ri.Get("/file_bucket/{guid}", file_bucket_handler.GetByGuid)
	})

	r.Get("/file_bucket/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/file_bucket", http.FileServer(http.Dir(config.Get().FileBucketPath))).ServeHTTP(w, r)
	})
	r.Get("/comfy_ui_gallery/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/comfy_ui_gallery", http.FileServer(http.Dir(config.Get().ComfyUIOutputDir))).ServeHTTP(w, r)
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
