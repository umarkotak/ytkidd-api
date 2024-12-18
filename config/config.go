package config

import (
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	config Config
)

func Initialize() {
	gotenv.Load()

	config = Config{
		AppEnv:            getEnvStringWithDefault("APP_ENV", "development"),
		AppPort:           getEnvStringWithDefault("APP_PORT", "33000"),
		AppHost:           getEnvStringWithDefault("APP_HOST", "http://localhost:33000"),
		DbURL:             os.Getenv("DB_URL"),
		DbTimezone:        getEnvStringWithDefault("DB_TIMEZONE", "Asia/Jakarta"),
		RedisUrl:          os.Getenv("REDIS_URL"),
		RedisPassword:     os.Getenv("REDIS_PASSWORD"),
		ChatWebsocketHost: getEnvStringWithDefault("CHAT_WEBSOCKET_HOST", "ws://localhost:33000/jodoh-express/api/chat/websocket/connect"),
		YoutubeApiKey:     os.Getenv("YOUTUBE_API_KEY"),
		FileBucketPath:    getEnvStringWithDefault("FILE_BUCKET_PATH", "file_bucket"),
		OllamaHost:        getEnvStringWithDefault("OLLAMA_HOST", "http://localhost:11434"),

		JxAuthJwtPrivateKey: os.Getenv("JX_AUTH_JWT_PRIVATE_KEY"),
		JxAuthJwtPublicKey:  os.Getenv("JX_AUTH_JWT_PUBLIC_KEY"),
		JxAuthJweSecretKey:  os.Getenv("JX_AUTH_JWE_SECRET_KEY"),

		ChatTokenJwtPrivateKey: os.Getenv("CHAT_TOKEN_JWT_PRIVATE_KEY"),
		ChatTokenJwtPublicKey:  os.Getenv("CHAT_TOKEN_JWT_PUBLIC_KEY"),

		InternalClientID:  os.Getenv("INTERNAL_CLIENT_ID"),
		InternalSecretKey: os.Getenv("INTERNAL_SECRET_KEY"),

		DevInternalClientID:  os.Getenv("DEV_INTERNAL_CLIENT_ID"),
		DevInternalSecretKey: os.Getenv("DEV_INTERNAL_SECRET_KEY"),
	}
}

func Get() Config {
	return config
}

func getEnvStringWithDefault(env, def string) string {
	if os.Getenv(env) == "" {
		return def
	}
	return os.Getenv(env)
}

func getEnvIntWithDefault(env string, def int) int {
	i, err := strconv.ParseInt(os.Getenv(env), 10, 64)
	if err != nil {
		return def
	}
	return int(i)
}
