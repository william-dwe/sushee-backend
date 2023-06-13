package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	ENV_MODE_TEST    = "test"
	ENV_MODE_RELEASE = "release"
)

type appConfig struct {
	Name string
	Url  string
}

type envConfig struct {
	Mode       string
	LoggerMode string
}

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type authConfig struct {
	TimeLimitAccessToken   string
	TimeLimitRefreshToken  string
	HmacSecretAccessToken  string
	HmacSecretRefreshToken string
}

type gcsConfig struct {
	ProjectID  string
	Bucket     string
	UploadPath string
}
type AppConfig struct {
	AppConfig  appConfig
	ENVConfig  envConfig
	DBConfig   dbConfig
	AuthConfig authConfig
	GCSConfig  gcsConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error loading .env file with these error details:\n%s", err))
	}
	log.Info().Msg(".env file Loaded")
}

func initConfig() AppConfig {
	initEnv()

	return AppConfig{
		AppConfig: appConfig{
			Name: getENV("APP_NAME", "sushee"),
			Url:  getENV("APP_URL", "127.0.0.1"),
		},
		ENVConfig: envConfig{
			Mode:       getENV("APP_ENV_MODE", ENV_MODE_TEST),
			LoggerMode: getENV("LOGGER_ENV_MODE", ENV_MODE_TEST),
		},
		DBConfig: dbConfig{
			Host:     getENV("CONF_DB_HOST", ""),
			User:     getENV("CONF_DB_USER", ""),
			Password: getENV("CONF_DB_PASSWORD", ""),
			DBName:   getENV("CONF_DB_NAME", ""),
			Port:     getENV("CONF_DB_PORT", ""),
		},
		AuthConfig: authConfig{
			TimeLimitAccessToken:   getENV("ACCESS_TOKEN_EXPIRATION", "900"),
			TimeLimitRefreshToken:  getENV("REFRESH_TOKEN_EXPIRATION", "86400"),
			HmacSecretAccessToken:  getENV("HMAC_SECRET_ACCESS_TOKEN", ""),
			HmacSecretRefreshToken: getENV("HMAC_SECRET_REFRESH_TOKEN", ""),
		},
		GCSConfig: gcsConfig{
			ProjectID:  getENV("GCS_PROJECT_ID", ""),
			Bucket:     getENV("GCS_BUCKET", ""),
			UploadPath: getENV("GCS_UPLOAD_PATH", ""),
		},
	}
}

var Config = initConfig()
