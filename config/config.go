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

type cloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
	Folder    string
}
type AppConfig struct {
	AppName          string
	ENVConfig        envConfig
	DBConfig         dbConfig
	AuthConfig       authConfig
	CloudinaryConfig cloudinaryConfig
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
		AppName: getENV("APP_NAME", "sushee"),
		ENVConfig: envConfig{
			Mode:       getENV("APP_ENV_MODE", ENV_MODE_TEST),
			LoggerMode: getENV("LOGGER_ENV_MODE", ENV_MODE_TEST),
		},
		DBConfig: dbConfig{
			Host:     getENV("DB_HOST", ""),
			User:     getENV("DB_USER", ""),
			Password: getENV("DB_PASSWORD", ""),
			DBName:   getENV("DB_NAME", ""),
			Port:     getENV("DB_PORT", ""),
		},
		AuthConfig: authConfig{
			TimeLimitAccessToken:   getENV("ACCESS_TOKEN_EXPIRATION", "900"),
			TimeLimitRefreshToken:  getENV("REFRESH_TOKEN_EXPIRATION", "86400"),
			HmacSecretAccessToken:  getENV("HMAC_SECRET_ACCESS_TOKEN", ""),
			HmacSecretRefreshToken: getENV("HMAC_SECRET_REFRESH_TOKEN", ""),
		},
		CloudinaryConfig: cloudinaryConfig{
			CloudName: getENV("CLOUDINARY_CLOUD_NAME", ""),
			APIKey:    getENV("CLOUDINARY_API_KEY", ""),
			APISecret: getENV("CLOUDINARY_API_SECRET", ""),
			Folder:    getENV("CLOUDINARY_PPROFILE_DIR", ""),
		},
	}
}

var Config = initConfig()
