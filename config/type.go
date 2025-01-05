package config

type (
	Config struct {
		AppEnv            string // Enum: development, integration, production
		AppPort           string //
		AppHost           string //
		DbURL             string //
		DbTimezone        string //
		RedisUrl          string //
		RedisPassword     string //
		ChatWebsocketHost string //
		YoutubeApiKey     string //
		FileBucketPath    string //
		OllamaHost        string //
		ComfyUIOutputDir  string

		CkAuthJwtPrivateKey string // jx auth - used for generating auth
		CkAuthJwtPublicKey  string // it will be jwt + jwe encryption
		CkAuthJweSecretKey  string //

		ChatTokenJwtPrivateKey string
		ChatTokenJwtPublicKey  string

		InternalClientID  string // used for other internal service when calling ytkidd-express-be service
		InternalSecretKey string //

		DevInternalClientID  string // used manually for dev for internal dev related API
		DevInternalSecretKey string //
	}
)
